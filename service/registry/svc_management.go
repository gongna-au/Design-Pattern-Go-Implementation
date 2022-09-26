package registry

import (
	"fmt"
	"github.com/Design-Pattern-Go-Implementation/db"
	"github.com/Design-Pattern-Go-Implementation/http"
	//"github.com/Design-Pattern-Go-Implementation/registry/model"
	"github.com/Design-Pattern-Go-Implementation/sidecar"
	//"github.com/google/uuid"
)

// svcManagement 服务管理，包含服务注册、更新、去注册。另外，服务订阅、去订阅、通知的功能由于与服务注册、更新、去注册紧密关联，
// 比如，每次的服务通知都是发生在服务状态变更之后，因此也把它们归到服务管理模块。

type svcManagement struct {
	localIp        string
	db             db.Db
	sidecarFactory sidecar.Factory
}

func newSvcManagement(localIp string, db db.Db, sidecarFactory sidecar.Factory) *svcManagement {
	return &svcManagement{
		localIp:        localIp,
		db:             db,
		sidecarFactory: sidecarFactory,
	}
}

// 服务注册
//（服务携带服务本身的信息以及这个服务想要被注册在那个注册中心的信息）
func (s *svcManagement) register(req *http.Request) *http.Response {

	// 根据这个服务的数据，知道了这个服务将要被注册在哪个注册中心中
	profile, ok := req.Body().(*model.ServiceProfile)
	if !ok {
		return http.ResponseOfId(req.ReqId()).
			AddStatusCode(http.StatusBadRequest).
			AddProblemDetails("service register request's body is not *ServiceProfile")
	}

	// 创建的事务名子往往是"register" + profile.Id
	// 注册有一系列的操作，这些操作往往都是要成功一起成功要失败一起失败
	transaction := s.db.CreateTransaction("register" + profile.Id)
	// 事务开始
	transaction.Begin()
	// Region 值对象，注册中心？更像是物理意义上的注册中心
	region := new(model.Region)
	// 去存储注册中心的表里去查具体的注册中心的信息（如果这个注册中心不存在需要插入）
	// 事务要执行的命令就是往注册中心表里面
	if err := s.db.Query(regionTable, profile.Region.Id, region); err != nil {
		cmd := db.NewInsertCmd(regionTable).WithPrimaryKey(profile.Region.Id).WithRecord(profile.Region)
		// transaction.Exec(cmd) 添加要执行的命令
		transaction.Exec(cmd)
	}
	// 然后往已经被注册的服务的信息表里面插入这个服务信息
	cmd := db.NewInsertCmd(profileTable).WithPrimaryKey(profile.Id).WithRecord(profile.ToTableRecord())
	transaction.Exec(cmd)

	if err := transaction.Commit(); err != nil {
		return http.ResponseOfId(req.ReqId()).AddStatusCode(http.StatusInternalServerError).
			AddProblemDetails(err.Error())
	}
	// 当某个服务被注册成功的时候我们需要通知给Register
	// 发送通知
	go s.notify(model.Register, profile)
	return http.ResponseOfId(req.ReqId()).AddStatusCode(http.StatusCreate)
}

// 注册中心表全部存放的注册中心的信息 注册中心1
// 服务更新
func (s *svcManagement) update(req *http.Request) *http.Response {
	profile, ok := req.Body().(*model.ServiceProfile)
	if !ok {
		return http.ResponseOfId(req.ReqId()).
			AddStatusCode(http.StatusBadRequest).
			AddProblemDetails("service update request's body is not *ServiceProfile")
	}
	transaction := s.db.CreateTransaction("register" + profile.Id)
	transaction.Begin()
	// 先更新regions表注册中心的信息
	// 更新的前提是表一定存在
	rcmd := db.NewUpdateCmd(regionTable).WithPrimaryKey(profile.Region.Id).WithRecord(profile.Region)
	transaction.Exec(rcmd)
	// 更新服务信息表
	pcmd := db.NewUpdateCmd(profileTable).WithPrimaryKey(profile.Id).WithRecord(profile.ToTableRecord())
	transaction.Exec(pcmd)
	if err := transaction.Commit(); err != nil {
		return http.ResponseOfId(req.ReqId()).AddStatusCode(http.StatusInternalServerError).
			AddProblemDetails(err.Error())
	}
	// 发送通知
	go s.notify(model.Update, profile)
	return http.ResponseOfId(req.ReqId()).AddStatusCode(http.StatusOk)
}

// 去注册
func (s *svcManagement) deregister(req *http.Request) *http.Response {
	// 去注册的时候只有一个服务ID
	svcId, ok := req.Header("service-id")
	if !ok {
		return http.ResponseOfId(req.ReqId()).
			AddStatusCode(http.StatusBadRequest).
			AddProblemDetails("service deregister request not contain service-id header")
	}
	// 根据服务ID先获取到完整的服务信息
	profileRecord := new(model.ServiceProfileRecord)
	var profile *model.ServiceProfile
	// 删除服务信息的前提是服务信息确实存在
	// 先查服务信息
	if err := s.db.Query(profileTable, svcId, profileRecord); err == nil {
		// 把服务信息转化为可以存储的服务信息
		profile = profileRecord.ToServiceProfile()
		region := new(model.Region)
		// 再先查注册中心在不在
		if err = s.db.Query(regionTable, profile.Region.Id, region); err != nil {
			return http.ResponseOfId(req.ReqId()).
				AddStatusCode(http.StatusInternalServerError).AddProblemDetails(err.Error())
		}
		profile.Region = region
		//注册中心在删除服务的的信息
		if err := s.db.Delete(profileTable, svcId); err != nil {
			return http.ResponseOfId(req.ReqId()).
				AddStatusCode(http.StatusInternalServerError).AddProblemDetails(err.Error())
		}
		// 发送通知
		go s.notify(model.Deregister, profile)
		return http.ResponseOfId(req.ReqId()).AddStatusCode(http.StatusNoContent)
	}
	return http.ResponseOfId(req.ReqId()).AddStatusCode(http.StatusBadRequest).
		AddProblemDetails("service-id " + svcId + " not exist")
}

// 服务订阅
func (s *svcManagement) subscribe(req *http.Request) *http.Response {
	subscription, ok := req.Body().(*model.Subscription)
	if !ok {
		return http.ResponseOfId(req.ReqId()).
			AddStatusCode(http.StatusBadRequest).
			AddProblemDetails("subscribe request's body is not Subscription")
	}
	if subscription.Id != "" {
		return http.ResponseOfId(req.ReqId()).
			AddStatusCode(http.StatusBadRequest).
			AddProblemDetails("Subscription Id is not empty")
	}
	subscription.Id = uuid.NewString()
	// 服务订阅就是往订阅表里插入一条订阅消息
	if err := s.db.Insert(subscriptionTable, subscription.Id, subscription); err != nil {
		return http.ResponseOfId(req.ReqId()).
			AddStatusCode(http.StatusInternalServerError).
			AddProblemDetails(err.Error())
	}
	return http.ResponseOfId(req.ReqId()).AddStatusCode(http.StatusCreate).
		AddHeader("subscription-id", subscription.Id)
}

// 服务去订阅
func (s *svcManagement) unsubscribe(req *http.Request) *http.Response {
	subscriptionId, ok := req.Header("subscription-id")
	if !ok {
		return http.ResponseOfId(req.ReqId()).
			AddStatusCode(http.StatusBadRequest).
			AddProblemDetails("service unsubscribe request not contain subscription-id header")
	}
	// 删除订阅信息
	if err := s.db.Delete(subscriptionTable, subscriptionId); err != nil {
		return http.ResponseOfId(req.ReqId()).
			AddStatusCode(http.StatusInternalServerError).AddProblemDetails(err.Error())
	}
	return http.ResponseOfId(req.ReqId()).AddStatusCode(http.StatusNoContent)
}

// 服务通知
func (s *svcManagement) notify(notifyType model.NotifyType, profile *model.ServiceProfile) {
	// 服务的ID 和服务的类型只要确定就产生了订阅这个服务的人？
	visitor := model.NewSubscriptionVisitor(profile.Id, profile.Type)
	// 在订阅表里面查，查到所有的订阅者，订阅者表的数据应该是 订阅者以及订阅的服务ID
	result, err := s.db.QueryByVisitor(subscriptionTable, visitor)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// 随即生成一个客户端
	// 在客户端
	httpClient, err := http.NewClient(s.sidecarFactory.Create(), s.localIp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, record := range result {
		// 遍历订阅者列表
		subscription := record.(*model.Subscription)
		// subscription 是具体的订阅者
		notification := model.NewNotification(subscription.Id)
		notification.Type = notifyType
		notification.Profile = profile.Clone().(*model.ServiceProfile)
		// 获取到具体的订阅者的URl
		notifyUri, err := subscription.NotifyUri()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		//获取到具体的订阅者的Endpoint
		notifyEndpoint, err := subscription.NotifyEndpoint()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		// 创建一个新的请求用来把控制通知的结构体notification)发往指定的URL
		req := http.EmptyRequest().AddUri(http.Uri(notifyUri)).AddMethod(http.POST).AddBody(notification)
		resp, err := httpClient.Send(notifyEndpoint, req)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("notify %s success, resp %+v", subscription.SrcSvcId, resp)
	}
}

package registry

import (
	"github.com/Design-Pattern-Go-Implementation/db"
	"github.com/Design-Pattern-Go-Implementation/http"
	"github.com/Design-Pattern-Go-Implementation/service/registry/model"
	"sort"
)

// svcDiscovery 服务发现
type svcDiscovery struct {
	db db.Db
}

func newSvcDiscovery(db db.Db) *svcDiscovery {
	return &svcDiscovery{db: db}
}

// 用来存储已经
type profiles []*model.ServiceProfile

// 服务发现
// 服务发现就是根据服务的类型的，以及服务唯一的id来查找可用的服务
// 服务管理侧重的是存储服务的信息 存储订阅者列表的信息 存储注册中心的信息 对这些信息的存储进行增删改查
func (s *svcDiscovery) discovery(req *http.Request) *http.Response {
	//根据服务的Id 服务的类型来查找服务
	svcId, _ := req.QueryParam("service-id")
	svcType, _ := req.QueryParam("service-type")
	// ServiceProfileVisitor profile表遍历, 筛选符合ServiceId和ServiceType的记录
	visitor := model.NewServiceProfileVisitor(svcId, model.ServiceType(svcType))
	//在对应的表里进行匹配
	result, err := s.db.QueryByVisitor(profileTable, visitor)
	if err != nil {
		return http.ResponseOfId(req.ReqId()).
			AddStatusCode(http.StatusInternalServerError).
			AddProblemDetails(err.Error())
	}
	profiles := make(profiles, 0)
	for _, record := range result { 
		//遍历匹配到的服务信息
		profile := record.(*model.ServiceProfileRecord).ToServiceProfile()
		// Region 值对象，每个服务都唯一属于一个Region
		// 在region表里面根据已有数据profile.Region.Id来查
		region := new(model.Region)
		if err := s.db.Query(regionTable, profile.Region.Id, region); err != nil {
			return http.ResponseOfId(req.ReqId()).
				AddStatusCode(http.StatusInternalServerError).
				AddProblemDetails(err.Error())
		}
		// 查找到的完整的region 重新赋值
		profile.Region = region
		// 然后才把数据完整的region存储
		profiles.add(profile)
	}
	// 优先返回优先级高的，如果优先级相等，则返回负载较小的
	sort.Sort(profiles)
	if len(profiles) == 0 {
		return http.ResponseOfId(req.ReqId()).
			AddStatusCode(http.StatusNotFound)
	}
	return http.ResponseOfId(req.ReqId()).AddStatusCode(http.StatusOk).AddBody(profiles[0])
}

/* 

func (m *memoryDb) QueryByVisitor(tableName string, visitor TableVisitor) ([]interface{}, error) {
	table, ok := m.tables.Load(tableName)
	if !ok {
		return nil, ErrTableNotExist
	}
	return table.(*Table).Accept(visitor)
}

func (t *Table) Accept(visitor TableVisitor) ([]interface{}, error) {
    return visitor.Visit(t)
}

func (s SubscriptionVisitor) Visit(table *db.Table) ([]interface{}, error) {
	var result []interface{}
	iter := table.Iterator()
	for iter.HasNext() {
		subscription := new(Subscription)
		if err := iter.Next(subscription); err != nil {
			return nil, err
		}
		// 先匹配ServiceId，如果一致则无须匹配ServiceType
		if subscription.TargetSvcId != "" && subscription.TargetSvcId == s.targetSvcId {
			result = append(result, subscription)
			continue
		}
		// ServiceId匹配不上，再匹配ServiceType
		if subscription.TargetSvcType != "" && subscription.TargetSvcType == s.targetSvcType {
			result = append(result, subscription)
		}
	}
	return result, nil
} */

package plugin

// Plugin 插件接口，在监控系统中，一切皆为插件
type Plugin interface {
	// Install 安装插件，只有在安装之后才能运行
	// 更准确的来说是让插件开始运行
	Install()
	// Uninstall 卸载插件，卸载后停止运行
	// 更准确的来说是让插件结束运行
	Uninstall()
	// SetContext 插件上下文设置 Context 是一个一直往下传的数据量不断增加的数据结构
	// 不断的存储不同的插件给予他的数据
	SetContext(ctx Context)
}

// 插件间需要传递一些数据？
type Event struct {
	headers map[string]string
	//playload any
	payload interface{}
}

func NewEvent(payload interface{}) *Event {
	return &Event{
		headers: make(map[string]string),
		payload: payload,
	}
}

func (e *Event) AddHeader(key, value string) *Event {
	e.headers[key] = value
	return e
}

func (e *Event) Payload() interface{} {
	return e.payload
}

func (e *Event) Header(key string) (string, bool) {
	val, ok := e.headers[key]
	return val, ok
}

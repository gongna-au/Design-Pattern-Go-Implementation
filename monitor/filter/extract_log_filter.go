package filter

import (
	"github.com/Design-Pattern-Go-Implementation/monitor/model"
	"github.com/Design-Pattern-Go-Implementation/monitor/plugin"
	"regexp"
)

// ExtractLogFilter 从日志中提 endpoint 和 model type
// 举例[192.168.1.1:8088][recv_req]receive request from address 192.168.1.91:80 success
// 则endpoint为192.168.1.1:8088，model type为recv_req
// ExtractLogFilter 具体的插件实现了Plugin 接口
type ExtractLogFilter struct {
	pattern *regexp.Regexp
}

func (e *ExtractLogFilter) Install() {
	e.pattern = regexp.MustCompile(`\[(.+)]\[(.+)].*`)
}

func (e *ExtractLogFilter) Uninstall() {
}

func (e *ExtractLogFilter) SetContext(ctx plugin.Context) {
}

// 具体的Filter动作
func (e *ExtractLogFilter) Filter(event *plugin.Event) *plugin.Event {
	log, ok := event.Payload().(string)
	if !ok {
		return event
	}
	//从log中匹配想要的字符串
	matches := e.pattern.FindStringSubmatch(log)
	if len(matches) != 3 {
		return event
	}
	re := model.NewMonitoryRecord()
	re.Endpoint = matches[1]
	re.Type = model.Type(matches[2])
	ev := plugin.NewEvent(re)
	return ev
}

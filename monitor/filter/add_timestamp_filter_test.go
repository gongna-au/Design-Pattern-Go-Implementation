package filter

import (
	"github.com/Design-Pattern-Go-Implementation/monitor/config"
	"github.com/Design-Pattern-Go-Implementation/monitor/model"
	"github.com/Design-Pattern-Go-Implementation/monitor/plugin"
	"testing"
)

func TestAddTimestampFilter(t *testing.T) {
	conf := config.Filter{
		Name:       "filter0",
		PluginType: "add_timestamp",
		Ctx:        plugin.EmptyContext(),
	}
	filterPlugin, err := NewPlugin(conf)
	if err != nil {
		t.Error(err)
	}
	filterPlugin.Install()
	re := model.NewMonitoryRecord()
	re.Endpoint = "192.168.0.1:80"
	re.Type = model.RecvResp
	event := plugin.NewEvent(re)
	event = filterPlugin.Filter(event)
	if event.Payload().(*model.MonitorRecord).Timestamp == 0 {
		t.Error("timestamp add failed")
	}
}

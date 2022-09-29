package input

import (
	"github.com/Design-Pattern-Go-Implementation/monitor/plugin"
	"github.com/Design-Pattern-Go-Implementation/mq"
)

// 内存消息中间件输入
type MemoryMqInput struct {
	//把获取的topic对应的数据存储起来
	topic mq.Topic
	//存储消费者实例
	consumer mq.Consumable
}

func (m *MemoryMqInput) Install() {
	//所有的具体插件一旦安装就意味着一定有这个插件对应的消费者的实例
	m.consumer = mq.MemoryMqInstance()
}

func (m *MemoryMqInput) Uninstall() {
}

func (m *MemoryMqInput) SetContext(ctx plugin.Context) {
	// 从Context map中获取到对应主题的数据
	// Context 插件配置上下文
	// type Context map[string]string
	if topic, ok := ctx.GetString("topic"); ok {
		//强制类型转化
		m.topic = mq.Topic(topic)
	}
}

func (m *MemoryMqInput) Input() (*plugin.Event, error) {
	msg, err := m.consumer.Consume(m.topic)
	if err != nil {
		return nil, err
	}

	event := plugin.NewEvent(msg.Payload())
	event.AddHeader("topic", string(m.topic))
	return event, nil
}

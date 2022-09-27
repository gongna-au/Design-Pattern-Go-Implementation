package mq

import (
	"errors"
	"sync"
)

/*
懒汉式单例模式，通过sync.Once实现
*/
var once = &sync.Once{}

// 定义一个memoryMq 内存消息队列，通过channel模式
// 作为一个消息队列它要接收很多类型的数据， 把同一个类型的数据放在一个topic对应的channel
type memoryMq struct {
	queues sync.Map // key为Topic，value为chan *Message，每个topic单独一个队列
}

//定义全局对象
var memoryMqInstance *memoryMq

func MemoryMqInstance() *memoryMq {
	once.Do(func() {
		memoryMqInstance = &memoryMq{
			queues: sync.Map{},
		}
	})
	return memoryMqInstance
}

func (m *memoryMq) Clear() {
	m.queues = sync.Map{}
}

func (m *memoryMq) Consume(topic Topic) (*Message, error) {
	messagechannel, ok := m.queues.Load(topic)
	if !ok {
		q := make(chan *Message, 10000)
		m.queues.Store(topic, q)
		messagechannel = q

	}
	queue, ok := messagechannel.(chan *Message)
	if !ok {
		return nil, errors.New("model's type is not chan *Message")
	}
	return <-queue, nil
}

func (m *memoryMq) Produce(message *Message) error {
	messagechannel, ok := m.queues.Load(message.topic)
	if !ok {
		q := make(chan *Message, 10000)
		m.queues.Store(message.topic, q)
		//对于新存储的数据，我们需要为它的topic新建立一个channel
		//这里值得注意的点是：channel 本质是指针 这里messagechannel 和q 都是指针
		//q指向一个已经分配好的内存地址
		//然后让messagechannel也指向这个分配好的内存地址
		//相当于messagechannel （这里是接口类型，其实应该是ch）和 q指向同一块内存 然后
		messagechannel = q
		return nil
	}
	//后期往ch写入数据就是往m.queues.Store(message.topic, q)里面的q写入数据
	ch, ok := messagechannel.(chan *Message)
	if !ok {
		return errors.New("model's type is not chan *Message")
	} else {
		ch <- message
		return nil
	}
}

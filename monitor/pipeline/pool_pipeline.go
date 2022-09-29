package pipeline

import (
	"fmt"
	"github.com/Design-Pattern-Go-Implementation/monitor/plugin"
	"github.com/panjf2000/ants"
)

// 创建一个大小为5的goroutine pool
// PoolPipeline 每次启动使用goroutine时启动
var pool, _ = ants.NewPool(5)

type PoolPipeline struct {
	pipelineTemplate
}

func (p *PoolPipeline) SetContext(ctx plugin.Context) {
	p.run = func() {
		// p有一个函数类型成员变量，用来存储结构体想要执行的行为
		if err := pool.Submit(p.doRun); err != nil {
			fmt.Printf("PoolPipeine run error %s", err.Error())
		}
	}
}

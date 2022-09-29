package pipeline

import "github.com/Design-Pattern-Go-Implementation/monitor/plugin"

// SimplePipeline 简单Pipeline实现，每次运行时新启一个goroutine
type SimplePipeline struct {
}

func (s *SimplePipeline) SetContext(ctx plugin.Context) {

}

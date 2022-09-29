package pipeline

import (
	"github.com/Design-Pattern-Go-Implementation/monitor/plugin"
	"reflect"
)

/*
 * 开闭原则（OCP）：一个软件系统应该具备良好的可扩展性，新增功能应当通过扩展的方式实现，而不是在已有的代码基础上修改
 * 根据具体的业务场景识别出那些最有可能变化的点，然后分离出去，抽象成稳定的接口。
 * 后续新增功能时，通过扩展接口，而不是修改已有代码实现
 * 例子：
 * pipeline.Plugin将输入、过滤、输出三个独立变化点，分离到三个接口input.Plugin、filter.Plugin、output.Plugin上，符合OCP
 */

/*
桥接模式
*/
type Plugin interface {
	plugin.Plugin
	SetInput(input input.Plugin)
	SetFilter(filter filter.Plugin)
	SetOutput(output output.Plugin)
}

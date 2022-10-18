package command

import "fmt"

// 负责发起请求。
// 此类必须具有用于存储对命令对象的引用的字段。
// 发送者触发该命令，而不是直接将请求发送给接收者。
// 请注意，发送者不负责创建命令对象。通常，它通过构造函数从客户端获取预先创建的命令。
type Invoker struct {
	command ICommand
}

func NewInvoker(i ICommand) *Invoker {
	return &Invoker{
		command: i,
	}
}

func (i *Invoker) SetCommand(ic ICommand) {
	i.command = ic
}

func (i *Invoker) ExcuteCommand() {
	//发送者触发该命令，而不是直接将请求发送给接收者
	i.command.Excute()
}

// 客户端创建和配置具体的命令对象。
// 客户端必须将所有请求参数（包括接收器实例）传递到命令的构造函数中。
type ICommand interface {
	Excute()
}

// 实现各种请求。
// 具体命令不应该自己执行工作，而是将调用传递给业务逻辑对象之一。
// 但是，为了简化代码，可以合并这些类。
// 命令类只处理请求如何传递给接收者的细节，而接收者自己做实际的工作。
type ConcreteCommand1 struct {
	receiver *Copy
	//receiver执行业务逻辑时需要的参数一定要被声明为具体命令的属性
	text string
}

func NewConcreteCommand1(c *Copy, text string) *ConcreteCommand1 {
	return &ConcreteCommand1{
		receiver: c,
		text:     text,
	}

}

func (c *ConcreteCommand1) SetReceiverParams(text string) {
	c.text = text
}

func (c *ConcreteCommand1) Excute() {
	c.receiver.CopyText(c.text)
}

// 命令类需要保存具体的业务receiver需要的参数
type ConcreteCommand2 struct {
	receiver *Paste
	//receiver执行业务逻辑时需要的参数一定要被声明为具体命令的属性
	text string
}

func NewConcreteCommand2(p *Paste, text string) *ConcreteCommand2 {
	return &ConcreteCommand2{
		receiver: p,
		text:     text,
	}
}

func (c *ConcreteCommand2) SetReceiverParams(text string) {
	c.text = text
}

// 命令类需要保存具体的业务receiver需要的参数或者只要把参数和receiver 一起作为参数传入调用就行
func (c *ConcreteCommand2) Excute() {
	c.receiver.PasteText(c.text)
}

// 具体的业务类函数
type Copy struct {
}

func NewCopy() *Copy {
	return &Copy{}
}

func (c *Copy) CopyText(text string) {
	fmt.Println(text + "text has copied successfully!")
}

type Paste struct {
}

func NewPaste() *Paste {
	return &Paste{}
}
func (c *Paste) PasteText(text string) {
	fmt.Println(text + "text has pasted successfully!")
}

func Client() {
	p := NewPaste()
	c := NewCopy()

	command1 := NewConcreteCommand1(c, "Hello")
	command2 := NewConcreteCommand2(p, "Hello")

	NewInvoker(command1).ExcuteCommand()
	NewInvoker(command2).ExcuteCommand()
}

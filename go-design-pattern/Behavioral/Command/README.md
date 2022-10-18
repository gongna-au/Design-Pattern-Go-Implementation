### 命令行为设计模式的结构：

#### 1.定义 ICommand 命令接口

```go
// 客户端创建和配置具体的命令对象。
// 客户端必须将所有请求参数（包括接收器实例）传递到命令的构造函数中。
type ICommand interface {
	Excute()
}

```

#### 2.定义业务逻辑接口或者具体的业务类

```go
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

```

#### 3.命令类向上"实现" 声明的命令接口，向下 "组合"业务逻辑接口『或者是具体的类』，（业务逻辑接口被称为接收者）

```go
type ConcreteCommand1 struct {
	receiver *Copy
	//receiver执行业务逻辑时需要的参数一定要被声明为具体命令的属性
	text string
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


func (c *ConcreteCommand2) SetReceiverParams(text string) {
	c.text = text
}

// 命令类需要保存具体的业务receiver需要的参数或者只要把参数和receiver 一起作为参数传入调用就行
func (c *ConcreteCommand2) Excute() {
	c.receiver.PasteText(c.text)
}

```

#### 4.业务逻辑 receiver 需要的参数是由具体的命令类传进入。（ 命令类『向下"组合"具体的业务逻辑』）所以命令类保存业务逻辑需要的参数

```go
func NewConcreteCommand1(c *Copy, text string) *ConcreteCommand1 {
	return &ConcreteCommand1{
		receiver: c,
		text:     text,
	}

}

func NewConcreteCommand2(p *Paste, text string) *ConcreteCommand2 {
	return &ConcreteCommand2{
		receiver: p,
		text:     text,
	}
}


```

#### 5.Invoker 组合 ICommand 命令接口，实现不同的 Invoker 类调用不同的实现了 ICommand 接口的命令实例，并在需要的时候调用。

```go
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

```

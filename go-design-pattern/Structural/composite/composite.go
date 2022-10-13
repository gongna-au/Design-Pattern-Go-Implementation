package composite

import (
	"errors"
	"fmt"
)

// 组合设计模式
// 只有当应用程序的核心模型可以表示为树时，使用组合模式才有意义
// 假设您决定创建一个使用这些类的订购系统。订单包含没有任何包装的简单产品，以及装满产品的盒子......和其他盒子。那么确定此类订单的总价格？
// 现实的做法：打开所有盒子，检查所有产品，然后计算总数。
// 程序的做法：必须事先知道具体的类，以及盒子的嵌套级别。才能调用具体的方法得到具体的价格。
type IComponent interface {
	Execute()
}

type Branches struct {
}

func NewBranches() *Branches {
	return &Branches{}
}

func (l *Branches) Execute() {
	fmt.Println("Branches")
}

type Leaf struct {
}

func NewLeaf() *Leaf {
	return &Leaf{}
}

func (l *Leaf) Execute() {
	fmt.Println("leaf")
}

type Composite struct {
	childs []IComponent
}

func NewComposite() *Composite {
	return &Composite{}
}

func (c *Composite) Execute() {
	for _, v := range c.childs {
		v.Execute()
	}

}

func (c *Composite) AddComposite(i IComponent) {
	c.childs = append(c.childs, i)

}

func (c *Composite) DeleteComposite(i int) error {

	if i < (len(c.childs)) && i > 0 {
		before := c.childs[:i-1]
		after := c.childs[i:]
		c.childs = append(before, after...)
		return nil
	}

	if len(c.childs) == 1 {
		c.childs = []IComponent{}
		return nil
	}
	if i >= len(c.childs) {
		return errors.New("index id not exists")
	}
	return nil

}

func (c *Composite) GetComposite(i int) (IComponent, error) {

	if i < (len(c.childs)) && i >= 0 {
		return c.childs[i], nil
	}
	if i >= len(c.childs) {
		return nil, nil
	}
	return nil, nil
}

func Client() {
	var c IComponent
	temp := NewComposite()
	temp.AddComposite(NewBranches())
	temp.AddComposite(NewLeaf())
	c = temp
	c.Execute()
}

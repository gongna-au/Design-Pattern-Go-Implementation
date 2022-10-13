### 代理模式的结构：

#### 1.IComponent 接口声明了不同的类他们共同实现的方法

```go
type IComponent interface {
	Execute()
}

```

#### 2.具体的 Leaf 类实现了接口的方法

```go
type Leaf struct {
}

func (l *Leaf) Execute() {

}

```

#### 3.具体的组合器——Composite 组合了 IComponent 切片

```go

type Composite struct {
	childs []IComponent
}

func (c *Composite) Execute() {

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


```

#### 4.客户端与接口直接

```go

func Client() {
	var c IComponent
	temp := NewComposite()
	temp.AddComposite(NewBranches())
	temp.AddComposite(NewComposite())
    c=temp
	c.Execute()
}

```

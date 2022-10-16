### 享元模式的结构：

#### 1.抽象出具体数据 TTree 对象之间可以共享的数据 TreeType，这里抽象出来的是树的状态，颜色，种类，高度

```go
type Tree struct {
	x             int
	y             int
	internalState TreeType
}

func NewTree(x int, y int, internalState TreeType) *Tree {

	return &Tree{
		x:             y,
		y:             y,
		internalState: internalState,
	}
}

func (t *Tree) Draw(x int, y int) {
	t.x = x
	t.y = y
}

```

#### 2. 因为数据对象的集合数量庞大消耗内存，所以我们才采取了享元模式，必然存在数据对象的集合 Forest.Trees ,一定存在让数据对象的集合里面继续添加数据元素的方法 AddTree，我更倾向与 AddTree（ t Tree）而不是 AddTree(x int, y int, name string, color string, texture string) ，使得 Forest 和 Tree 解耦

```go

type Forest struct {
	Trees []Tree
}

func NewForest() *Forest {
	return &Forest{
		Trees: []Tree{},
	}
}

func (t *Forest) AddTree(x int, y int, name string, color string, texture string) {
	result := treeTypeFactory.GetTreeType(name, color, texture)
	tree := NewTree(x, y, result)
	t.Trees = append(t.Trees, *tree)
}
```

#### 3.对于抽象用于共享的数据出来的享元数据对象

```go
// 在此示例中，享元模式有助于在画布上渲染数百万个树对象时减少内存使用。
type TreeType struct {
	name    string
	color   string
	texture string
}

func NewTreeType(name string, color string, texture string) TreeType {
	return TreeType{
		name:    name,
		color:   color,
		texture: texture,
	}
}

// 享元对象本身可以改变自己的状态
func (t *TreeType) ChangeName(name string) {
	t.name = name
}

func (t *TreeType) ChangeColor(color string) {
	t.name = color
}

func (t *TreeType) ChangeTexture(text string) {
	t.texture = text
}

```

#### 4.抽象出享元工厂，享元数据工厂对享元数据对象的集合进行操作，供给客户端进行调用。

```go

// 享元工厂保存享元对象
type TreeTypeFactory struct {
	Trees *TreeTypeCollection
}

func NewTreeTypeFactory(t ...TreeType) *TreeTypeFactory {
	trees := []TreeType{}
	trees = append(trees, t...)
	result := TreeTypeCollection(trees)
	return &TreeTypeFactory{
		Trees: &result,
	}
}

// 封装得到享元的复杂过程
func (t *TreeTypeFactory) GetTreeType(name string, color string, texture string) TreeType {
	result, err := t.Trees.Find(name, color, texture)
	if err != nil {
		newTreeType := NewTreeType(name, color, texture)
		t.Trees.Add(newTreeType)
		return newTreeType
	} else {
		return result
	}
}

// 得到所有的享元
func (t *TreeTypeFactory) GetAllTrees() []TreeType {
	return *t.Trees

}

// 输出所有的享元
func (t *TreeTypeFactory) OutputAllTrees() {
	for k, v := range *t.Trees {
		fmt.Printf("The " + strconv.Itoa(k) + " is")
		fmt.Println(v)
	}
}

// Find 是供给GetTreeType()调用的
func (t TreeTypeCollection) Find(name string, color string, texture string) (TreeType, error) {
	for _, v := range t {
		if v.name == name && v.color == color && v.texture == texture {
			return v, nil
		} else {
			continue
		}

	}
	return TreeType{}, errors.New("not find")
}

// Add 是供给GetTreeType()调用的
func (t *TreeTypeCollection) Add(trees ...TreeType) {
	*t = append([]TreeType(*t), trees...)
}

```

package flyweight

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
)

// 声明并且初始化
var once = sync.Once{}

// 只声明
var treeTypeFactory *TreeTypeFactory

func init() {
	if treeTypeFactory == nil {
		once.Do(func() {
			treeTypeFactory = NewTreeTypeFactory()
		})
	}

}

// 主 Tree 类中提取重复的内在状态并将其移动到享元类 TreeType 中
// 我们希望从某个对象中提取出来这个对象的状态
// 1.不是把相同的数据存储在多个对象中，
// 2.而是将其保存在几个享元对象中，并链接到充当上下文的适当 Tree 对象。
// 3.享元工厂保存享元对象的集合，并且封装(从集合中得到享元对象)(往集合中添加新的享元对象)的复杂过程。（客户端代码使用享元工厂）
// 4.
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

// Forest 代表很多Tree 的集合，因为每个Tree保存的状态大多差不多
// 且对于Forest来说，它拥有的Tree太多了，而且每个Tree存储需要的内存很大
// 所以考虑把Tree中的internalState 抽离出来构成享元
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

type TreeTypeCollection []TreeType

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

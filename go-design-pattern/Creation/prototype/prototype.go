package prototype

import (
	"encoding/json"
	"time"
)

/*
这个模式在 Java、C++ 这种面向对象的语言不太常用，
但是如果大家使用过 javascript 的话就会非常熟悉了，
因为 js 本身是基于原型的面向对象语言，所以原型模式在 js 中应用非常广泛。
接下来会按照一个类似课程中的例子使用深拷贝和浅拷贝结合的方式进行实现
需求:
假设现在数据库中有大量数据，
包含了关键词，关键词被搜索的次数等信息，
模块 A 为了业务需要会在启动时加载这部分数据到内存中
并且需要定时更新里面的数据
同时展示给用户的数据每次必须要是相同版本的数据，
不能一部分数据来自版本 1 一部分来自版本 2
*/

// Keyword 搜索关键字
type Keyword struct {
	word      string
	visit     int
	UpdatedAt *time.Time
}

// Clone 这里使用序列化与反序列化的方式深拷贝
func (k *Keyword) Clone() *Keyword {
	var newKeyword Keyword
	// 把要拷贝的结构体json.Marshal 进行序列化
	// 然后在反序列化json.Unmarshal 把序列化好的数据和要返回的数据的指针传进去
	// 返回指针即可
	b, _ := json.Marshal(k)
	json.Unmarshal(b, &newKeyword)
	return &newKeyword
}

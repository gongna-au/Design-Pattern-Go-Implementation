package builder

import "fmt"

// 建造者模式：构造函数、set 方法、建造者模式三种对象创建方式
// 其实在 Golang 中对于创建类参数比较多的对象的时候，我们常见的做法是必填参数直接传递，可选参数通过传递可变的方法进行创建。
// Go 编写建造者模式的代码其实会很长
// 这些是它的一个缺点，所以如果不是参数的校验逻辑很复杂的情况下一般我们在 Go 中不会采用这种方式，
// 而会采用后面的另外一种方式

const (
	defaultMaxTotal = 10
	defaultMaxIdle  = 9
	defaultMinIdle  = 1
)

// ResourcePoolConfig resource pool
type ResourcePoolConfig struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

// ResourcePoolConfigBuilder 用于构建 ResourcePoolConfig
type ResourcePoolConfigBuilder struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

// SetName SetName
func (b *ResourcePoolConfigBuilder) SetName(name string) error {
	if name == "" {
		return fmt.Errorf("name can not be empty")
	}
	b.name = name
	return nil
}

// SetMinIdle SetMinIdle
func (b *ResourcePoolConfigBuilder) SetMinIdle(minIdle int) error {
	if minIdle < 0 {
		return fmt.Errorf("max tatal cannot < 0, input: %d", minIdle)
	}
	b.minIdle = minIdle
	return nil
}

// SetMaxIdle SetMaxIdle
func (b *ResourcePoolConfigBuilder) SetMaxIdle(maxIdle int) error {
	if maxIdle < 0 {
		return fmt.Errorf("max tatal cannot < 0, input: %d", maxIdle)
	}
	b.maxIdle = maxIdle
	return nil
}

// SetMaxTotal SetMaxTotal
func (b *ResourcePoolConfigBuilder) SetMaxTotal(maxTotal int) error {
	if maxTotal <= 0 {
		return fmt.Errorf("max tatal cannot <= 0, input: %d", maxTotal)
	}
	b.maxTotal = maxTotal
	return nil
}

// 建造者和被建造者之间的关系
// Build Build
func (b *ResourcePoolConfigBuilder) Build() (*ResourcePoolConfig, error) {
	if b.name == "" {
		return nil, fmt.Errorf("name can not be empty")
	}

	// 设置默认值
	if b.minIdle == 0 {
		b.minIdle = defaultMinIdle
	}

	if b.maxIdle == 0 {
		b.maxIdle = defaultMaxIdle
	}

	if b.maxTotal == 0 {
		b.maxTotal = defaultMaxTotal
	}

	if b.maxTotal < b.maxIdle {
		return nil, fmt.Errorf("max total(%d) cannot < max idle(%d)", b.maxTotal, b.maxIdle)
	}

	if b.minIdle > b.maxIdle {
		return nil, fmt.Errorf("max idle(%d) cannot < min idle(%d)", b.maxIdle, b.minIdle)
	}
	// 用建造者携带的参数去赋值给被建造者
	return &ResourcePoolConfig{
		name:     b.name,
		maxTotal: b.maxTotal,
		maxIdle:  b.maxIdle,
		minIdle:  b.minIdle,
	}, nil
}

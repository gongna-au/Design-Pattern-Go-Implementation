### 桥接模式的结构

#### 1.抽象部分 （Abstrac­tion） 提供高层控制逻辑， 依赖于完成底层实际工作的实现对象。

```go
// Remote类则作为抽象 这部分实现了一些复杂的行为，这些行为依赖具体的实现部分的操作
// “抽象部分”定义了两个类层次结构中“控制”部分的接口。它管理着一个指向“实
// 现部分”层次结构中对象的引用，并会将所有真实工作委派给该对象。
type RemoteControl struct {
	device Device
}

// 暴露给客户端调用
func NewRemoteControl(device Device) *RemoteControl {
	return &RemoteControl{
		device: device,
	}
}
func (r *RemoteControl) TogglePower() error {
	if r.device.isEnable() {
		r.device.disable()
		return nil
	} else {
		r.device.enable()
		return nil
	}
}

func (r *RemoteControl) VolumeDown() {
	r.device.setVolume(r.device.getVolume() - 10)
}

func (r *RemoteControl) VolumeUp() {
	r.device.setVolume(r.device.getVolume() + 10)
}

func (r *RemoteControl) ChannelDown() {
	r.device.setVolume(r.device.getChannel() - 10)
}

func (r *RemoteControl) ChannelUp() {
	r.device.setVolume(r.device.getChannel() + 10)

}



```

#### 2.实现部分 （Imple­men­ta­tion） 为所有具体实现声明通用接口。 抽象部分仅能通过在这里声明的方法与实现对象交互。抽象部分可以列出和实现部分一样的方法， 但是抽象部分通常声明一些复杂行为， 这些行为依赖于多种（被实现部分声明）的原语操作。

```go
type Device interface {
	isEnable() bool
	enable()
	disable()
	getVolume() int
	setVolume(int)
	getChannel() int
	setChannel(int)
}


```

#### 3.具体实现 （Con­crete Imple­men­ta­tions） 中包括特定于平台的代码。

```go
// 被拆分的类
type Tv struct {
}

func NewTv() *Tv {
	return &Tv{}
}

func (t *Tv) isEnable() bool {
	return true
}
func (t *Tv) enable() {

}
func (t *Tv) disable() {

}
func (t *Tv) getVolume() int {
	return 10
}
func (t *Tv) setVolume(int) {

}
func (t *Tv) getChannel() int {
	return 10

}
func (t *Tv) setChannel(int) {

}

// 被拆分的类
type Radio struct {
}

func NewRadio() *Radio {
	return &Radio{}
}
func (t *Radio) isEnable() bool {
	return true
}
func (t *Radio) enable() {

}
func (t *Radio) disable() {

}
func (t *Radio) getVolume() int {
	return 100

}
func (t *Radio) setVolume(int) {

}
func (t *Radio) getChannel() int {
	return 100
}
func (t *Radio) setChannel(int) {

}

```

#### 4.精确抽象 （Refined Abstrac­tion） 提供控制逻辑的变体。 与其父类一样， 它们通过通用实现接口与不同的实现进行交互。

```go
type AdvancedRemote struct {
	RemoteControl
}

func (a *AdvancedRemote) Mute() {
	a.device.setVolume(0)

}

```

#### 5.通常情况下， 客户端 （Client） 仅关心如何与抽象部分合作。 但是， 客户端需要将抽象对象与一个实现对象连接起来。

```go
    remote := NewRemoteControl(NewTv())
	err := remote.TogglePower()
	require.Nil(t, err)
```

### 桥接模式的结构

#### 1.抽象的 RemoteControl 组合了 Device 接口 ，RemoteControl 就是一个“桥“，外界通过这个“桥“，调用这个 Device 接口下的方法

```go
// Remote类则作为抽象 这部分实现了一些复杂的行为，这些行为依赖具体的实现部分的操作
// “抽象部分”定义了两个类层次结构中“控制”部分的接口。它管理着一个指向“实
// 现部分”层次结构中对象的引用，并会将所有真实工作委派给该对象。
// 暴露给子类使用调用
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

#### 2.Device 接口

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

#### 3.Tv 、Radio 具体实现了 Device 接口

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

#### 4.AdvancedRemote 具体的桥，继承了 RemoteControl ，便于通过横向扩展具体的桥达到扩展 RemoteControl 抽象的桥的功能。

```go
type AdvancedRemote struct {
	*RemoteControl
}

func (a *AdvancedRemote) Mute() {
	a.device.setVolume(0)

}
func NewAdvancedRemote(d Device) *AdvancedRemote {
	return &AdvancedRemote{
		RemoteControl: NewRemoteControl(d) ,
	}
}


```

#### 5.通常情况下， 客户端 （Client） 仅关心如何与抽象部分合作。 但是， 客户端需要将抽象对象与一个实现对象连接起来。

```go
func Client(){
	remote := NewAdvancedRemote(NewRemoteControl(NewTv()))
	err := remote.TogglePower()
	require.Nil(t, err)
}

```

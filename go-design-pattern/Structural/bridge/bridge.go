package bridge

import "fmt"

/* 1. 处理继承的时候，试图在两个独立的维度（两个父类）扩展子类的时候：
子类可以同时继承这两个维度的类，
也可以是抽取其中一个维度的类并使之成为独立的类层次，那么在初始类中引用这个新层次的对象，从而使得一个类不必拥有所有的状态和行为
“把一个类层次转化为多个相关的类层次”
*/
// 设备Device类作为实现（不同于编程中的实现概念）部分， 而 遥控器Remote类则作为抽象部分。
type Device interface {
	isEnable() bool
	enable()
	disable()
	getVolume() int
	setVolume(int)
	getChannel() int
	setChannel(int)
}

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

// Remote类则作为抽象 这部分实现了一些复杂的行为，这些行为依赖具体的实现部分的操作
// “抽象部分”定义了两个类层次结构中“控制”部分的接口。它管理着一个指向“实
// 现部分”层次结构中对象的引用，并会将所有真实工作委派给该对象。
type RemoteControl struct {
	device Device
}

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

type AdvancedRemote struct {
	*RemoteControl
}

func (a *AdvancedRemote) Mute() {
	a.device.setVolume(0)

}
func NewAdvancedRemote(d Device) *AdvancedRemote {
	return &AdvancedRemote{
		RemoteControl: NewRemoteControl(d),
	}
}

func Client() {
	remote := NewAdvancedRemote(NewTv())
	err := remote.TogglePower()
	if err != nil {
		fmt.Println(err)
	}

}

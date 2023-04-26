package common

import (
	"fmt"
)

// 接口定义
type Usb interface {
	Start()
	Stop()
	// 私有方法，只能在本包中使用
	privatePrint()
}

// Mp3 将接口两个方法都实现了
type Mp3 struct {
	Name string
}

func (c *Mp3) Start() {
	fmt.Println(c.Name, "start working")
}
func (c *Mp3) Stop() {
	fmt.Println(c.Name, "stop work")
}
func (c *Mp3) privatePrint() {
	fmt.Println("private print")
}

// Phone 只实现了start方法
type Phone struct {
	Name string
}

func (p *Phone) Start() {
	fmt.Println(p.Name, "start working")
}

// 定义使用者的结构体
type Computer struct {
	Name string
}

// Working方法接收Usb接口对象，然后调用它的两个方法
func (c *Computer) Working(u Usb) {
	u.Start()
	u.Stop()
	u.privatePrint()
}

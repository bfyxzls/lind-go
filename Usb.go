package main

import (
	"fmt"
	"lind-go/common"
	_ "lind-go/common" //自定义的包
)

// 接口定义
type Usb interface {
	Start()
	Stop()
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
}

func main() {
	fmt.Println("auth:%s", common.Name)
	mp3 := Mp3{"mp3"}
	computer := Computer{"computer"}
	computer.Working(&mp3)

	//phone := Phone{"phone"} //没有实现usb的两个方法
	//computer.Working(&phone) //报错
}

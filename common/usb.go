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

// 大写表示公有方法
func (c *Mp3) Stop() {
	fmt.Println(c.Name, "stop work")
}

// 小写表示私有方法
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
	// computer的一个属性
	Name string
	Subs map[int]*Computer
}

// (c *Computer)表示这个方法是Computer类型的
// Working方法表示是computer类型的方法；Usb表示一个入参对象，在方法体里调用Usb的三个方法start,stop和privatePrint
func (c *Computer) Working(u Usb) {
	u.Start()
	u.Stop()
	u.privatePrint()
	c.Subs = map[int]*Computer{}
	c.Subs[1] = &Computer{Name: "1"}
	c.Subs[2] = &Computer{Name: "2"}
	c.Subs[3] = &Computer{Name: "3"}
	delete(c.Subs, 2)
	for i := range c.Subs {
		fmt.Println("i=", i)
	}
	defer fmt.Println("c.name=", c.Name)
	// go func() {} 表示创建一个新的 Goroutine(轻量级线程),用于异步执行函数。
	go func() {
		fmt.Println("hello world")
	}()

	// defer是Go语言提供的一种用于注册延迟调用的机制：让函数或语句可以在当前函数执行完毕后（包括通过return正常结束或者panic导致的异常结束）执行。

}

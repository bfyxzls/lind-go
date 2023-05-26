package main

import (
	"fmt"
	"lind-go/common"
	"regexp"
	"time"
)

func main() {
	fmt.Printf("auth:%s", common.Name)
	mp3 := common.Mp3{"mp3"} // :=这个运算符可以使变量在不声明的情况下直接被赋值使用
	computer := common.Computer{"computer", nil}
	computer.Working(&mp3)
	//phone := Phone{"phone"} //没有实现usb的两个方法
	//computer.Working(&phone) //报错

	// 时间测试
	now := time.Now()
	// 返回日期
	year, month, day := now.Date()
	fmt.Printf("year:%d, month:%d, day:%d\n", year, month, day)
	// 格式化
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("15:04:05"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	// 时间戳转换
	layout := "2006-01-02 15:04:05"
	t := time.Unix(now.Unix(), 0) // 参数分别是：秒数,纳秒数
	fmt.Println(t.Format(layout))
	const text = "My email is ccmouse@gmail.com"
	// 正则表达式
	compile := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]+`)
	match := compile.FindString(text)
	fmt.Println(match)

	// 直接初始化切片slices
	slices := []int{1, 2, 3}
	slices[1] = 22
	fmt.Println(slices)
}

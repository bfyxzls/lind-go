package main

import (
	"fmt"
	"lind-go/common"
)

func main() {
	fmt.Printf("auth:%s", common.Name)
	mp3 := common.Mp3{"mp3"}
	computer := common.Computer{"computer"}
	computer.Working(&mp3)

	//phone := Phone{"phone"} //没有实现usb的两个方法
	//computer.Working(&phone) //报错
}

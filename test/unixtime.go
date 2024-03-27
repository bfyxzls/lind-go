package main

import (
	"fmt"
	"strconv"
	"time"
)

// 测试时间戳
func main() {

	for i := 1; i < 120; i++ {
		now := time.Now()
		minuteAligned := now.Truncate(time.Minute)
		timeStamp := strconv.FormatInt(minuteAligned.Unix(), 10)
		fmt.Println(now.String() + "..." + timeStamp)
		time.Sleep(time.Second)
	}

}

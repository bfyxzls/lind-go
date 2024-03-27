package util

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// 测试时间戳
func TestSleep(t *testing.T) {
	now := time.Now()
	minuteAligned := now.Truncate(time.Minute)
	timeStamp := strconv.FormatInt(minuteAligned.Unix(), 10)

	for i := 1; i < 120; i++ {
		fmt.Println(time.Now().String() + "..." + timeStamp)
		time.Sleep(time.Second)
	}
}
func TestTimeOutFormat(t *testing.T) {
	// 5秒后超时
	now := time.Now()
	formattedDate := now.Format("20060102030405")
	fmt.Println(formattedDate)
}

package util

import (
	"fmt"
	"testing"
	"time"
)

func TestTotp(t *testing.T) {
	for i := 1; i < 10; i++ {
		fmt.Println(GenerateTOTP([]byte("pkulaw"), 30))
		time.Sleep(5 * time.Second) // 每隔5秒再测试一次
	}
}

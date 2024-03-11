package util

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	result := BinarySearch([]string{"risklens-test.pkulaw.cn"}, "aliyun-test-lawyer-mgr.pkulaw.com")
	fmt.Println(result)
}

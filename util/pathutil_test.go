package util

import (
	"fmt"
	"testing"
)

func TestPathRegex(t *testing.T) {
	// 创建一个路径过滤器
	pathFilter := NewPathFilter("/swagger-ui/*")

	// 测试路径匹配
	fmt.Println(pathFilter.Matches("/swagger-ui/abc"))          // true
	fmt.Println(pathFilter.Matches("/swagger-ui/abc/bcd"))      // true
	fmt.Println(pathFilter.Matches("/path/swagger-ui/abc/bcd")) // true
	fmt.Println(pathFilter.Matches("/api/docs"))                // false

	// 创建一个路径过滤器
	pathFilter2 := NewPathFilter("/swagger-ui")
	fmt.Println(pathFilter2.Matches("/swagger-ui")) // false

}

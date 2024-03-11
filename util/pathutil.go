package util

import (
	"fmt"
	"regexp"
)

// 定义路径过滤器结构体
type PathFilter struct {
	Pattern string
}

// 初始化路径过滤器
func NewPathFilter(pattern string) *PathFilter {
	return &PathFilter{Pattern: pattern}
}

// 路径匹配方法
func (pf *PathFilter) Matches(path string) bool {
	matched, err := regexp.MatchString(pf.Pattern, path)
	if err != nil {
		fmt.Printf("路径匹配出错：%v\n", err)
		return false
	}
	return matched
}

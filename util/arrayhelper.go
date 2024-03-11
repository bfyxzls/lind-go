package util

// 从数组中查找元素 BinarySearch 二分查找
func BinarySearch(strs []string, target string) bool {
	for _, s := range strs {
		if s == target {
			return true
		}
	}
	return false
}

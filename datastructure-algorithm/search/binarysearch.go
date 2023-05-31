package search

// BinarySearch 二分查找算法
func BinarySearch(a []int, target int) int {
	i, j := 0, len(a)
	for i <= j {
		m := (i + j) / 2
		if a[m] == target {
			// 找到了
			return m
		} else if a[m] > target {
			j = m - 1
		} else {
			i = m + 1
		}
	}
	// 没有找到
	return -1
}

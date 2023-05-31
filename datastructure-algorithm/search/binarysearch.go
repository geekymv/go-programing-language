package search

// BinarySearch 二分查找算法
func BinarySearch(a []int, target int) int {
	i, j := 0, len(a)-1
	for i <= j {
		m := (i + j) / 2
		if target == a[m] {
			// 找到了
			return m
		} else if target < a[m] {
			j = m - 1
		} else { // target > a[m]
			i = m + 1
		}
	}
	// 没有找到
	return -1
}

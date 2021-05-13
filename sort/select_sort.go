package sort

/*
选择排序
时间复杂度：O(n^2)
空间复杂度：O(1)
*/

func SelectSort(a []int, n int) {
	if n <= 1 {
		return
	}
	for i := 0; i < n; i++ {
		// 查找最小值
		minIndex := i
		for j := i + 1; j < n; j++ {
			if a[j] < a[minIndex] {
				minIndex = j
			}
		}
		// 交换
		a[i], a[minIndex] = a[minIndex], a[i]
	}
}

package sort

/*
插入排序
时间复杂度：O(n^2)
空间复杂度：O(1)
*/

func InsertSort(a []int, n int) {
	if n <= 1 {
		return
	}
	for i := 1; i < n; i++ {
		val := a[i] // 表示要插入的数
		j := i - 1
		for ; j >= 0; j-- {
			if a[j] > val {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		a[j] = val
	}
}

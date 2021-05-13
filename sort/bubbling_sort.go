package sort

/*
冒泡排序
时间复杂度: O(n^2)
空间复杂度: O(1)
*/
func BubbleSort(a []int, n int) {
	for i := 0; i < n; i++ {
		var flag bool // 提前退出冒泡循环的标志位
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				flag = true // 表示有数据交换
			}
		}
		if !flag {
			return // 没有数据交换，提前退出
		}
	}
}

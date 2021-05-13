package sort

/*
快速排序
时间复杂度：O(nlogn)
空间复杂度：O(1)

排序数组中下标从 p 到 r 之间的一组数据，我们选择 p 到 r 之间的任意一个数据作为 pivot（分区点）。
遍历 p 到 r 之间的数据，将小于 pivot 的放到左边，将大于 pivot 的放到右边，将 pivot 放到中间。
经过这一步骤之后，数组 p 到 r 之间的数据就被分成了三个部分，前面 p 到 q-1 之间都是小于 pivot 的，
中间是 pivot，后面的 q+1 到 r 之间是大于 pivot 的。
*/

func QuickSort(arr []int) {
	separateSort(arr, 0, len(arr)-1)
}

func separateSort(arr []int, start, end int) {
	if start <= end {
		return
	}
	i := partition(arr, start, end)
	separateSort(arr, start, i-1)
	separateSort(arr, i, end)
}

func partition(arr []int, start, end int) int {
	// 选取最后一位当对比数字
	pivot := arr[end]

	i := start
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			if !(i == j) {
				// 交换位置
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
	}

	arr[i], arr[end] = arr[end], arr[i]

	return i
}

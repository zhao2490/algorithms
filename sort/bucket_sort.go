package sort

import "fmt"

/*
桶排序
时间复杂度：O(n)
空间复杂度：O()
*/

func getMax(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func BucketSortSimple(source []int) {
	if len(source) <= 1 {
		return
	}
	array := make([]int, getMax(source)+1)
	for i := 0; i < len(source); i++ {
		array[source[i]] ++
	}
	fmt.Println(array)
	c := make([]int, 0)
	for i := 0; i < len(array); i++ {
		for array[i] != 0 {
			c = append(c, i)
			array[i]--
		}
	}
	copy(source, c)
}

func BucketSort(source []int) {
	num := len(source)
	if num <= 1 {
		return
	}
	max := getMax(source)
	buckets := make([][]int, num)

	index := 0
	for i := 0; i < num; i++ {
		index = source[i] * (num - 1) / max // 桶序号
		buckets[index] = append(buckets[index], source[i])
	}

	tmpPos := 0 // 标记数组位置
	for i := 0; i < num; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0 {
			QuickSort(buckets[i]) // 桶内做快排
			copy(source[tmpPos:], buckets[i])
			tmpPos += bucketLen
		}
	}
}

/*
计数排序
计数排序其实是桶排序的一种特殊情况。当要排序的 n 个数据，所处的范围并不大的时候，比如最大值是 k，我们就可以把数据划分成 k 个桶。每个桶内的数据值都是相同的，省掉了桶内排序的时间。
时间复杂度：O(n)
空间复杂度：O(n)
*/

func CountingSort(arr []int, n int) {
	if n <= 1 {
		return
	}
	// 查找数组中数据的范围
	max := arr[0]
	for i := 1; i < n; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	c := make([]int, max+1) // 申请一个计数数组c，[0,max]int
	// 计算每个元素的个数
	for i := 0; i < n; i++ {
		c[arr[i]]++
	}

	// 依次累加
	for i := 1; i <= max; i++ {
		c[i] = c[i-1] + c[i]
	}

	// 临时数组r，存储排序之后的结果
	r := make([]int, n)
	// 计数排序关键步骤
	for i := n - 1; i >= 0; i-- {
		index := c[arr[i]] - 1
		r[index] = arr[i]
		c[arr[i]]--
	}
	// 将结果拷贝给arr数组
	for i := 0; i < n; i++ {
		arr[i] = r[i]
	}
}

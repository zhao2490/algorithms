package sort

import (
	"testing"
)

func TestInsertSort(t *testing.T) {
	array := []int{2, 5, 7, 3, 1, 2}
	InsertSort(array, len(array))
	t.Log(array)
}

func TestBubbleSort(t *testing.T) {
	array := []int{2, 5, 7, 3, 1, 2}
	BubbleSort(array, len(array))
	t.Log(array)
}

func TestSelectSort(t *testing.T) {
	array := []int{2, 5, 7, 3, 1, 2}
	SelectSort(array, len(array))
	t.Log(array)
}

func TestQuickSort(t *testing.T) {
	array := []int{2, 5, 7, 3, 1, 2}
	QuickSort(array)
	t.Log(array)
}

func TestMergeSort(t *testing.T) {
	array := []int{2, 5, 7, 3, 1, 2}
	MergeSort(array)
	t.Log(array)
}

func TestBucketSort(t *testing.T) {
	array := []int{2, 5, 7, 3, 1, 2}
	BucketSort(array)
	t.Log(array)
}

package LinkedList

import (
	"strconv"
	"testing"
)

func Test_LRUHashMap(t *testing.T) {
	lruHashMap := NewLRUHashMap(3)
	for i := 0; i < 6; i++ {
		lruHashMap.Set(strconv.Itoa(i), i)
	}
	val := lruHashMap.Get("1")
	t.Logf("Get 1: %v", val)
	val = lruHashMap.Get("4")
	t.Logf("Get 4: %v", val)
	lruHashMap.Del("4")
	h := lruHashMap.linkedList.Head
	for h.Next != nil {
		t.Log(h.Data)
		h = h.Next
	}
}

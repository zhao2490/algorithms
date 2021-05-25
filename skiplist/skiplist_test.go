package skiplist

import "testing"

func TestSkipList(t *testing.T) {
	sl := NewSkipList()

	sl.Insert("zhao", 95)
	t.Log(sl.head.forwards[0])
	t.Log(sl.head.forwards[0].forwards[0])
	t.Log(sl)
	t.Log("-----------------------------")

	sl.Insert("leo", 88)
	t.Log(sl)
	t.Log("-----------------------------")
	sl.Insert("jack", 100)
	t.Log(sl)
	t.Log("-----------------------------")

	t.Log(sl.Find("jack", 100))
	t.Log("-----------------------------")

	t.Log(sl.Delete("jack", 100))
	t.Log(sl.Find("jack", 100))
	t.Log("-----------------------------")

	sl.Delete("leo", 95)
	t.Log(sl)
	t.Log("-----------------------------")

}

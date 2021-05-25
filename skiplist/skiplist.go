package skiplist

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	MAX_LEVEL = 16 // 最大层数
)

type skipListNode struct {
	v        interface{}     // 跳表保存的值
	score    int             // 用于排序分数
	level    int             // 层高
	forwards []*skipListNode // 每层前进的指针
}

// 新建跳表节点
func newSkipListNode(v interface{}, score, level int) *skipListNode {
	return &skipListNode{v: v, score: score, level: level, forwards: make([]*skipListNode, level, level)}
}

// 跳表结构体
type SkipList struct {
	head   *skipListNode // 头节点
	level  int           // 跳表当前层数
	length int           // 跳表长度
}

// 实例化跳表对象
func NewSkipList() *SkipList {
	// 头节点，便于操作
	head := newSkipListNode(0, math.MinInt32, MAX_LEVEL)
	return &SkipList{head: head, level: 1, length: 0}
}

func (sl *SkipList) Length() int {
	return sl.length
}

func (sl *SkipList) Insert(v interface{}, score int) int {
	if v == nil {
		return 1
	}
	// 查找插入位置
	cur := sl.head
	// 记录每层路径
	update := [MAX_LEVEL]*skipListNode{}
	i := MAX_LEVEL - 1
	for ; i >= 0; i-- {
		for cur.forwards[i] != nil {
			if cur.forwards[i].v == v {
				return 2
			}
			if cur.forwards[i].score > score {
				// 记录所有level需要修改的节点
				update[i] = cur
				break
			}
			cur = cur.forwards[i]
		}
		if cur.forwards[i] == nil {
			update[i] = cur
		}
	}

	// 通过随机算法获取该节点层数
	level := 1
	for i := 1; i < MAX_LEVEL; i++ {
		if rand.Int31()%7 == 1 {
			level++
		}
	}

	// 创建一个新的跳表节点
	newNode := newSkipListNode(v, score, level)

	// 原有节点连接 通过取出随机level，修改索引
	for i := 0; i <= level-1; i++ {
		next := update[i].forwards[i]
		update[i].forwards[i] = newNode
		newNode.forwards[i] = next
	}

	// 如果当前节点的层数大于之前跳表的层数
	// 更新当前跳表层数
	if level > sl.level {
		sl.level = level
	}

	// 更新跳表长度
	sl.length++
	return 0
}

// 查找
func (sl *SkipList) Find(v interface{}, score int) *skipListNode {
	if v == nil || sl.length == 0 {
		return nil
	}

	cur := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for cur.forwards[i] != nil {
			if cur.forwards[i].score == score && cur.forwards[i].v == v {
				return cur.forwards[i]
			} else if cur.forwards[i].score > score {
				// 当前节点 score < 目标score 向下查找
				break
			}
			// cur.forwards[i].score < score 向后遍历
			cur = cur.forwards[i]
		}
	}

	return nil
}

// 删除节点
func (sl *SkipList) Delete(v interface{}, score int) int {
	if v == nil {
		return 1
	}

	// 查找前驱节点
	cur := sl.head
	// 记录前驱路径
	update := [MAX_LEVEL]*skipListNode{}
	for i := sl.level - 1; i >= 0; i-- {
		update[i] = sl.head
		for cur.forwards[i] != nil {
			if cur.forwards[i].score == score && cur.forwards[i].v == v {
				update[i] = cur
				break
			}
			cur = cur.forwards[i]
		}
	}
	// 执行删除操作
	cur = update[0].forwards[0]
	for i := cur.level - 1; i >= 0; i-- {
		if update[i] == sl.head && cur.forwards[i] == nil {
			sl.level = i
		}

		if update[i].forwards[i] == nil {
			update[i].forwards[i] = nil
		} else {
			update[i].forwards[i] = update[i].forwards[i].forwards[i]
		}
	}

	sl.length--
	return 0
}

func (sl *SkipList) String() string {
	return fmt.Sprintf("level:%+v, length:%+v", sl.level, sl.length)
}

package LinkedList

// 使用HashMap和双向链表实现LRU缓存
type LRUHashMap struct {
	hashMap    map[string]*LinkedNode
	linkedList *LinkedList
}

type LinkedNode struct {
	Key  string
	Data interface{}
	Next *LinkedNode
	Prev *LinkedNode
}

type LinkedList struct {
	Head     *LinkedNode // 头节点
	Length   int         // 当前长度
	Capacity int         // 最大容量
}

func NewLRUHashMap(capacity int) *LRUHashMap {
	linkedList := &LinkedList{
		Capacity: capacity,
		Head:     &LinkedNode{},
	}
	return &LRUHashMap{
		hashMap:    make(map[string]*LinkedNode),
		linkedList: linkedList,
	}
}

// 删除链表中最后一个元素
func (l *LinkedList) deletedElementEnd() *LinkedNode {
	p := l.Head
	if p.Next == nil {
		return p
	}
	for p.Next.Next != nil {
		p = p.Next
	}
	end := p.Next
	end.Prev.Next = nil
	l.Length--
	return end
}

// 在链表头部插入元素
func (l *LinkedList) insertElementHead(node *LinkedNode) {
	if l.Length == 0 {
		l.Head = node
		l.Length++
		return
	}
	node.Next = l.Head
	node.Next.Prev = node
	l.Head = node
	l.Length++
}

// 删除node
func (l *LinkedList) deleteElement(node *LinkedNode) {
	node.Prev.Next = node.Next
	l.Length--
}

func (m *LRUHashMap) Set(key string, value interface{}) {
	node := &LinkedNode{
		Key:  key,
		Data: value,
		Next: nil,
		Prev: nil,
	}
	if oldNode, ok := m.hashMap[key]; ok {
		m.linkedList.deleteElement(oldNode)
	}
	m.linkedList.insertElementHead(node)
	m.hashMap[key] = node
	if m.linkedList.Length > m.linkedList.Capacity {
		// 超出容量时删除链表最末尾节点
		endNode := m.linkedList.deletedElementEnd()
		delete(m.hashMap, endNode.Key)
	}
}

func (m *LRUHashMap) Get(key string) interface{} {
	node, ok := m.hashMap[key]
	if ok {
		m.linkedList.deleteElement(node)
		m.linkedList.insertElementHead(node)
		return node.Data
	}
	return nil
}

func (m *LRUHashMap) Del(key string) {
	node, ok := m.hashMap[key]
	if ok {
		m.linkedList.deleteElement(node)
		delete(m.hashMap, key)
	}
}

package trie_tree

/*
	仅支持26个字母 a-z
	如果想优化 扩大容纳字符 可以使用有序数组 跳表 或者红黑树 记录子节点
	适用于 输入法自动补全，输入网址自动补全，热门搜索关键词等场景
*/
type TrieTreeNode struct {
	data         rune
	children     [26]*TrieTreeNode
	isEndingChar bool
}

type TrieTree struct {
	root *TrieTreeNode
}

// 在TrieTree中插入
func (t *TrieTree) Insert(text []rune) {
	p := t.root
	for i := 0; i < len(text); i++ {
		index := text[i] - 'a'
		if p.children[index] == nil {
			newNode := &TrieTreeNode{
				data:         text[i],
				children:     [26]*TrieTreeNode{},
				isEndingChar: false,
			}
			p.children[index] = newNode
		}
		p = p.children[index]
	}
	p.isEndingChar = true
}

// 在TrieTree中查找
func (t *TrieTree) Find(pattern []rune) bool {
	p := t.root
	for i := 0; i < len(pattern); i++ {
		index := pattern[i] - 'a'
		if p.children[index] == nil {
			return false
		}
		p = p.children[index]
	}
	return p.isEndingChar
}


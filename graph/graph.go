package graph

import (
	"container/list"
	"fmt"
)

type Graph struct {
	adj []*list.List
	v   int
}

func newGraph(capacity int) Graph {
	g := Graph{
		adj: make([]*list.List, capacity+1),
		v:   capacity,
	}
	for i := range g.adj {
		g.adj[i] = list.New()
	}
	return g
}

func (g *Graph) addEdge(s int, t int) {
	g.adj[s].PushBack(t)
	g.adj[t].PushBack(s)
}

// 广度优先
func (g *Graph) BFS(s, t int) {
	if s == t {
		return
	}
	// prev 用来记录搜索路径。当我们从顶点 s 开始，广度优先搜索到顶点 t 后，prev 数组中存储的就是搜索的路径。不过，这个路径是反向存储的。prev[w]存储的是，顶点 w 是从哪个前驱顶点遍历过来的
	prev := make([]int, g.v)
	for i := 0; i < g.v; i++ {
		prev[i] = -1
	}
	// queue 用来存储已经被访问、但相连的顶点还没有被访问的顶点
	var queue []int
	queue = append(queue, s)
	// visited 是用来记录已经被访问的顶点，用来避免顶点被重复访问。如果顶点 q 被访问，那相应的 visited[q]会被设置为 true
	visited := make([]bool, g.v)
	visited[s] = true
	isFound := false
	for len(queue) > 0 && !isFound {
		top := queue[0]
		queue = queue[1:]
		linkedList := g.adj[top]
		for e := linkedList.Front(); e != nil; e = e.Next() {
			k := e.Value.(int)
			if !visited[k] {
				prev[k] = top
				if k == t {
					isFound = true
					break
				}
				queue = append(queue, k)
				visited[k] = true
			}
		}
	}

	if isFound {
		printPrev(prev, s, t)
	} else {
		fmt.Printf("no path found from %d to %d\n", s, t)
	}
}

func printPrev(prev []int, s, t int) {
	if t == s || prev[t] == -1 {
		fmt.Printf("%d", t)
	} else {
		printPrev(prev, s, prev[t])
		fmt.Printf("%d", t)
	}
}

// 深度优先
func (g *Graph) DFS(s, t int) {

	prev := make([]int, g.v)
	for i := range prev {
		prev[i] = -1
	}

	visited := make([]bool, g.v)
	visited[s] = true

	isFound := false
	g.recurDFS(s, t, visited, prev, &isFound)

	if isFound {
		printPrev(prev, s, t)
	} else {
		fmt.Printf("no path found from %d to %d\n", s, t)
	}
}

func (g *Graph) recurDFS(s, t int, visited []bool, prev []int, isFound *bool) {

	if *isFound {
		return
	}

	visited[s] = true

	if s == t {
		*isFound = true
		return
	}

	linkedList := g.adj[s]
	for e := linkedList.Front(); e != nil; e = e.Next() {
		k := e.Value.(int)
		if !visited[k] {
			prev[k] = s
			g.recurDFS(k, t, visited, prev, isFound)
		}
	}
}

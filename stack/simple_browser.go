package stack

import "fmt"

/*
使用栈实现浏览器页面前进 回退
*/
type Browser struct {
	nowAddr      string
	forwardStack *Stack
	backStack    *Stack
}

func NewBrowser() *Browser {
	return &Browser{
		forwardStack: NewStack(),
		backStack:    NewStack(),
	}
}

func (b *Browser) CanForward() bool {
	if b.forwardStack.IsEmpty() {
		return false
	}
	return true
}

func (b *Browser) CanBack() bool {
	if b.backStack.IsEmpty() {
		return false
	}
	return true
}

func (b *Browser) Open(addr string) {
	fmt.Printf("open new addr %+v\n", addr)
	b.forwardStack.Flush()
	if b.nowAddr != "" {
		b.PushBack(b.nowAddr)
	}
	b.nowAddr = addr
}

func (b *Browser) PushBack(addr string) {
	b.backStack.Push(addr)
}

func (b *Browser) Forward() {
	if b.forwardStack.IsEmpty() {
		return
	}
	top := b.forwardStack.Pop()
	b.backStack.Push(b.nowAddr)
	b.nowAddr = top.(string)
	fmt.Printf("forward to %+v\n", top)
}

func (b *Browser) Back() {
	if b.backStack.IsEmpty() {
		return
	}
	top := b.backStack.Pop()
	b.forwardStack.Push(b.nowAddr)
	b.nowAddr = top.(string)
	fmt.Printf("back to %+v\n", top)
}

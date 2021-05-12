package stack

import "testing"

func Test_SimpleBrowser(t *testing.T) {
	b := NewBrowser()
	b.Open("www.qq.com")
	b.Open("www.baidu.com")
	b.Open("www.taobao.com")

	if b.CanBack() {
		b.Back()
	}
	if b.CanForward() {
		b.Forward()
	}
	if b.CanBack() {
		b.Back()
	}
	if b.CanBack() {
		b.Back()
	}
	if b.CanBack() {
		b.Back()
	}
	b.Open("www.sina.com")
	if b.CanForward() {
		b.Forward()
	}
	if b.CanBack() {
		b.Back()
	}
}

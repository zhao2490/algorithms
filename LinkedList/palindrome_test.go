package LinkedList

import "testing"

func TestPalindrome1(t *testing.T) {
	strs := []string{"heooeh", "hello", "heoeh", "a", ""}
	for _, str1 := range strs {
		l := NewSingleLinkedList(str1)
		t.Log(isPalindrome1(l))
	}
}

func TestPalindrome2(t *testing.T) {
	strs := []string{"heooeh", "hello", "heoeh", "a", ""}
	for _, str1 := range strs {
		l := NewSingleLinkedList(str1)
		t.Log(isPalindrome2(l))
	}
}

func TestPalindrome3(t *testing.T) {
	strs := []string{"heooeh", "hello", "heoeh", "a", ""}
	for _, str1 := range strs {
		l := NewSingleLinkedList(str1)
		t.Log(isPalindrome3(l))
	}
}

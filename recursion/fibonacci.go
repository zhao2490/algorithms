package recursion

import "fmt"

/*
斐波那契
*/

type Fibs struct {
	val map[int]int
}

func NewFibs(n int) *Fibs {
	return &Fibs{
		val: make(map[int]int, n),
	}
}

func (fib *Fibs) Fibonacci(n int) int {
	if v, ok := fib.val[n]; ok {
		return v
	}
	if n <= 1 {
		fib.val[1] = 1
		return 1
	} else if n == 2 {
		fib.val[2] = 1
		return 1
	} else {
		res := fib.Fibonacci(n-2) + fib.Fibonacci(n-1)
		fib.val[n] = res
		return res
	}
}

func (fib *Fibs) Print(n int) {
	fmt.Println(fib.val[n])
}

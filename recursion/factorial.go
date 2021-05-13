package recursion

import "fmt"

/*
迭代实现阶乘
*/
type Fac struct {
	Val map[int]int
}

func NewFactorial(n int) *Fac {
	return &Fac{make(map[int]int, n)}
}

func (fac *Fac) Factorial(n int) int {
	if v, ok := fac.Val[n]; ok {
		return v
	}
	if n <= 1 {
		fac.Val[1] = 1
		return 1
	} else {
		res := n * fac.Factorial(n-1)
		fac.Val[n] = res
		return res
	}
}

func (fac *Fac) Print(n int) {
	fmt.Println(fac.Val[n])
}

package test

import "fmt"

type A struct {
	d int
}

func (a A) F2(x int) {
	a.d += x
}
func Testmain() {
	a := &A{d: 10}
	a.F2(5)
	fmt.Println(*a)
}

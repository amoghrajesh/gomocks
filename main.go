package main

import "fmt"

type myInterface interface {
	foo(i int, s string)
}

type realImpl struct {
	x int
}

func (r realImpl) foo(i int, s string) {
	fmt.Printf("The parameters I have are int %d and string %s \n", i, s)
	r.x = 5
	fmt.Printf("The value of x is %d\n", r.x)
}

func main() {
	r := realImpl{x: 0}
	r.foo(1, "hello")
}

package main

import "fmt"

type reaImpl struct{}

func (r reaImpl) IsUnder16(i int) bool {
	fmt.Println("Value of i is", i)
	return i < 16
}

func main() {
	fmt.Println("Inside main")
}

package main

import "fmt"

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}

func foo(stuff ...int) {
	for _, n := range stuff {
		fmt.Printf("%v ", n)
	}
	fmt.Printf("\n")
}

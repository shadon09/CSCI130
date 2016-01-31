package main

import "fmt"

func main() {
	half := func(x int) string {
		y := x / 2
		even := x%2 == 0
		return fmt.Sprint(y, even)
	}
	fmt.Printf("1: %v\n", half(1))
	fmt.Printf("2: %v\n", half(2))
}

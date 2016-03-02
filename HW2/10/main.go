package main

import "fmt"

func main() {
	n := []int{5, 43, 54, 2, 9}
	fmt.Printf("Slice: %v\n", n)
	k := greatest(n...)
	fmt.Printf("Greatest Value: %v\n", k)
}

func greatest(k ...int) int {
	s := 0
	for _, v := range k {
		if s < v {
			s = v
		}
	}
	return s
}

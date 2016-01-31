package main

import "fmt"

func main() {
	var a int
	var b int
	var c int
	fmt.Println("Enter a number: ")
	fmt.Scan(&a)
	fmt.Println("Enter a bigger number: ")
	fmt.Scan(&b)
	c = b % a
	fmt.Printf("The remainder of %v/%v is %v \n", b, a, c)
}

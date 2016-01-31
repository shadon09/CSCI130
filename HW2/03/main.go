package main

import "fmt"

func main() {
	var name string
	fmt.Printf("Please enter first name \n")
	fmt.Scan(&name)
	fmt.Printf("Hello %v \n", name)
}

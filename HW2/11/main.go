package main

import "fmt"

func main() {
	n := true && false || false && true || !(false && false)
	fmt.Printf("%v\n", n)
}

//Answer is True

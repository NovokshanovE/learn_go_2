package main

import (
	"fmt"
)

func test() (x int) {
	defer func() {
		fmt.Printf("x1=%d\n", x)

	}()
	defer func() {
		x++
	}()
	defer func() {
		fmt.Printf("x2=%d\n", x)

	}()

	x = 1
	return
}

func anotherTest() int {
	var x int
	defer fmt.Println(x)
	defer func() {
		fmt.Printf("x1=%d\n", x)
		x++

	}()

	x = 1
	return x
}

func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}

package main

import (
	"fmt"
)

func invalidIfCondition() {
	var x int = 10
	var y int = 20
	/*
		if x > y {
			fmt.Printf("%d is greater than %d\n", x, y);
		}
		else if x < y {
			fmt.Printf("%d is greater than %d\n", y, x);
		}
		else {
			fmt.Printf("%d is equal %d\n", x, y);
		}
	*/

	if x > y {
		fmt.Printf("%d is greater than %d\n", x, y)
	} else if x < y {
		fmt.Printf("%d is greater than %d\n", y, x)
	} else {
		fmt.Printf("%d is equal %d\n", x, y)
	}
}

func assignAndIf() bool {
	var valX int = 50
	if valY := 30; valX > valY {
		fmt.Printf("%d is greater than %d\n", valX, valY)
		return true
	}
	return false
}

func returnCondition() bool {
	var c, y int = 10, 20
	return c == y
}

func main() {
	var a, c int = 5, 1

	if a > c {
		// F for format like %d
		fmt.Printf("%d is greater than %d", a, c)
	} else if c > a {
		fmt.Printf("%d is greater than %d", c, a)
	} else {
		fmt.Printf("%d is equal", a)
	}

	fmt.Printf("\n")
	invalidIfCondition()
	fmt.Println(assignAndIf())
	fmt.Println(returnCondition())
}

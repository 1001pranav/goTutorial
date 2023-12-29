package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println("Hello, World!");

	// Variables and data.
	// Shorthand operator for assigning variables.
	message := "Assignment Operator";

	// Integer declarations and multiple declarations
	var number int;

	// Multiple value assignment
	var num1, num2 int = 1, 2;

	fmt.Println(message, number, num1, num2);

	fmt.Println("Using the function from other module.. ",rand.Intn(120));

	const CONSTANT_STRING string = "Hello ಜಗತ್ತು";
	fmt.Println(CONSTANT_STRING);
}
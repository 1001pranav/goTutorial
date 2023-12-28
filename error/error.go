package main

import (
	"fmt"
	"errors"
)
func recoverPanicExample(dividend int, divisor int) int {
	/*
		defer is commonly used for tasks such as closing files,
		releasing resources, or recovering from panics.
		It helps keep cleanup and setup code close together and improves code readability.
	*/
	defer func () {
		fmt.Println("Closing in the end");
		/*
			Recover is used to regain the control after panic,
			it will return nil if not exists
		*/
		if r := recover(); r != nil {
			fmt.Println("Recovered: ", r);
		}
	}()

	if (divisor == 0) {
		panic("Divisor cannot be 0");
	}
	return dividend/divisor;
}

func divideTwoNumbers(dividend int, divisor int) (int, error){
	if divisor == 0 {
		return 0, errors.New("Divisor cannot be zero");
	}

	return dividend/divisor, nil;
}

func main() {
	quotient, err := divideTwoNumbers(25, 5);
	if (err != nil) {
		fmt.Println("Error is -", err);
	} else {
		fmt.Println("Quotient is -", quotient);
	}

	quotient, err = divideTwoNumbers(25, 0);
	if (err != nil) {
		fmt.Println("Error is -", err);
	} else {
		fmt.Println("Quotient is -", quotient);
	}

	recoverPanicExample(12,0);
}
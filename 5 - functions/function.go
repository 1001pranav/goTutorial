package main

import "fmt"

/*
	* func keyword is used to define new function.
	* if we want to take any parameters to the function then we can use (paramName <dataType>)
		else we can keep it empty ()
	* if function return any data then after
		`func dec() <return DataType> {}`
		* To return arrayDataType
			`func <functionName>() <[]DataType> {...}`
	needs to be specified.
	* If a function has multiple return parameters then we can use following template
		`func dec() (<return DataType>, <return DataType>) {}`
	* Just like spread Operator in JavaScript in golang we can use variadic function
		`func variadic(v ...<DataType>){}` v will have array data
	* we can create an anonymous function which acts as a closure.
	* Named return values
*/

func addTwoNumbers(num1, num2 int) int {
	return num1 + num2
}

/* Variadic function */
func addNumbers(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

/* Closure and Anonymous function */
func counter() (func() int, func() int) {
	var count int = 0

	increment := func() int {
		count++
		return count
	}

	decrement := func() int {
		count--
		return count
	}

	return increment, decrement
}

func splitNumbers(number int) (splitNumber1, splitNumber2 int) {
	splitNumber1 = number / 2
	splitNumber2 = number / 2

	/*
		This is also called naked return because variableName is already declared in return type.
		Here the values of the variables declared in the function declarations are returned when return keyword is called.
	*/

	return
}

func main() {

	increment, decrement := counter()

	fmt.Println("Incrementing...,", increment())
	fmt.Println("Incrementing...,", increment())
	fmt.Println("Decrementing...,", decrement())
	fmt.Println("Incrementing...,", increment())
	fmt.Println("Incrementing...,", increment())
	fmt.Println("Decrementing...,", decrement())
	fmt.Println("Incrementing...,", increment())
	fmt.Println("Decrementing...,", decrement())

	numbers := []int{1, 2, 3, 4, 5}
	total := addNumbers(numbers...)
	fmt.Println(total)
	total = addNumbers(total, 25, 30, 40, 50)
	fmt.Println(total)

	fmt.Println(addTwoNumbers(34, 25))
	fmt.Println(splitNumbers(total))
}

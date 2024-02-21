package main

import "fmt"

func arrayAssignment() {
	// Arrays are fixed in the size where we can't change once the array's are initialized
	var arr [3]int

	// Pre declare arrays
	var arr1 = [3]int{5, 5, 25}

	// Shorthand for arrays
	arr2 := [8]int{3, 5, 9, 11, 26, 21, 223}

	// initialization with ellipses
	a := [...]int{25, 5, 5} //complier determines length
	// Different sizes of arrays cannot be assigned.

	// MultiDimensional arrays
	var multiDimArr = [2][2]int{
		{1, 2},
		{2, 3}, //Trailing comma is mandatory
	}

	subArray := [5]int{2: 3, 4, 5}

	// Pass by value;
	passByValue := arr2

	// Pass by Reference;
	passByRef := &arr2

	passByRef[0] = 12

	fmt.Println("Passing by value", passByValue)               //Passing by value [3 5 9 11 26 21 223 0]
	fmt.Println("Pass By reference", *passByRef)               //Pass By reference [12 5 9 11 26 21 223 0]
	fmt.Println("Array after updating referenced Array", arr2) // Array after updating referenced Array [12 5 9 11 26 21 223 0]

	fmt.Println("Add array while initialization from position ", subArray) //Add array while initialization from position  [0 0 3 4 5]
	// Array comparison with "=="
	// Arrays in both variables should be in same place and same value.

	fmt.Println(a == arr)  //false
	fmt.Println(a == arr1) //false

	fmt.Println("Multi-dimensional arrays", multiDimArr) //Multi-dimensional arrays [[1 2] [2 3]]
	fmt.Println(arr2)                                    //[12 5 9 11 26 21 223 0]

	fmt.Println("Length of the array is -> ", len(arr2)) //Length of the array is ->  8

	arr3 := [8]int{3, 5, 9, 11, 26, 21, 223}

	// Array[startIndex(inclusive): endIndex(exclusive)]
	split := arr3[2:5]
	fmt.Println(split) // [9 11 26]

	slices := []int{}
	fmt.Println(slices)   // []
	fmt.Println(arr3[:5]) //[3, 5, 9, 11, 26]
}

func sliceExample() {
	// Array
	arr := [8]int{3, 5, 9, 11, 26, 21, 223}

	//Creating slices from array
	sliceArr := arr[:len(arr)-1]
	//Slices are the references of the array, Here slice is the reference of the array
	fmt.Println("sliceArr is  ", sliceArr, arr)

	slicesArr1 := arr[:]
	fmt.Println(slicesArr1)
	// Slices are the data types which holds similar data without the limit
	// Creating slices using make
	slice := make([]int, 5, 10)
	fmt.Println(slice) //[0 0 0 0 0]
	/*
		* In the above example complier will allocate 10 memory.
		* Here from index 0 to 4 compiler will be free to assign the default data
		* From 5 to 9 compiler will be create empty spaces. We need use append method add data.
		* EG [0,0,0,0,0,*,*,*,*]
			so length of above slice is 5 and capacity is 10.
	*/

	// To add an element to the slice we use append method
	slice = append(slice, 255) // This will add element in the end
	fmt.Println(slice)

	cp := make([]int, len(slice))

	// This will copy the values only wont reference the slice
	copy(cp, slice)

	cp = append(cp, 500)
	cp[1] = 25
	fmt.Println(cp)
	fmt.Println(slice)

	/*
		- The length of a slice is the number of elements it contains.

		- The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

		- The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).
	*/
	fmt.Println("length of slice", len(slice), "Capacity:", cap(slice))

	/*
		- An empty slice is considered as 'nil'
		- This has length and capacity as 0.
	*/
	var emptySlice []int
	fmt.Println(emptySlice == nil, len(emptySlice), cap(emptySlice), emptySlice)
}

func sliceLiterals() {
	sl := []struct {
		s string
		i int
	}{
		{"John", 4},
		{"Alice", 5},
		{"Bob", 3},
		{"Doe", 3},
	}

	fmt.Println(sl)
}

func forLoops() {
	arr := [8]int{3, 5, 9, 11, 26, 21, 223}

	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			fmt.Println("Empty value in the array", i)
		} else {
			fmt.Println("Value of array is ", arr[i])
		}
	}

	i := 0
	for i < len(arr) {
		arr[i] += i
		i++
	}

	i = 0
	for {
		fmt.Println("Array value is: ", arr[i])
		i++
		if i >= len(arr) {
			break
		}
	}

	for index, arrVal := range arr {
		fmt.Println("Index is ", index, " Value of the array is", arrVal)
	}
}

func main() {
	arrayAssignment()
	sliceExample()
	forLoops()
	sliceLiterals()
}

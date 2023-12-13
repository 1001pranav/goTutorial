package main

import "fmt"

func arrayAssignment() {
		var arr [3] int;

		// Pre declare arrays
		var arr1 = [3]int{5, 5, 25};

		// Shorthand for arrays
		arr2 := [8]int{3,5,9,11,26, 21, 223};

		// initialization with ellipses
		a := [...]int{25,5,5}; //complier determines length
		// Different sizes of arrays cannot be assigned.

		// MultiDimensional arrays
		var multiDimArr = [2][2]int{
			{1,2},
			{2,3},//Trailing comma is mandatory
		};

		subArray := [5]int{2:3,4,5};

		// Pass by value;
		passByValue := arr2;

		// Pass by Reference;
		passByRef := &arr2;

		passByRef[0] = 12;

		fmt.Println("Passing by value", passByValue);
		fmt.Println("Pass By reference", *passByRef);
		fmt.Println("Array after updating referenced Array", arr2);

		fmt.Println("Add array while initialization from position ", subArray);
		// Array comparison with "=="
		// Arrays in both variables should be in same place and same value.

		fmt.Println(a == arr);
		fmt.Println(a == arr1);

		fmt.Println("Multi-dimensional arrays", multiDimArr);
		fmt.Println(arr2);

}

func sliceExample() {

	arr2 := [8]int{3,5,9,11,26, 21, 223};

	// Array[startIndex(inclusive): endIndex(exclusive)]
	slices := arr2[2:5];
	fmt.Println(slices);

	slices2 := []int{}
	fmt.Println(slices2);
	// Creating slices using make
	slice3 := make([]int, 5, 10);
	fmt.Println(slice3);
	/*
		* In the above example complier will allocate 10 memory.
		* Here from index 0 to 4 compiler will be free to assign the default data
		* From 5 to 9 compiler will be create empty spaces. We need use append method add data.
		* EG [0,0,0,0,0,*,*,*,*]
			so length of above slice is 5 and capacity is 10.

	*/
}


func main() {
	arrayAssignment();
	sliceExample();
}
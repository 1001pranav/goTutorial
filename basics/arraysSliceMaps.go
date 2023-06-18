package main

import "fmt"

func main() {
	// Arrays 
	var arr [3] int;

	// Pre declare arrays
	var arr1 = [3]int{5, 5, 25};
	
	// Shorthand for arrays
	arr2 := [8]int{3,5,9,11,26, 21, 223};
	
	a := [...]int{25,5,5}; //complier determines length
	// Different sizes of arrays cannot be assigned.

	// MultiDimensional arrays
	var multiDimArr = [2][2]int{
		{1,2},
		{2,3},//Trailing comma is mandatory
	}

	// Array comparison with "=="
	// Arrays in both variables should be in same place and same value.

	fmt.Println(a==arr);
	fmt.Println(a==arr1);
	fmt.Println(multiDimArr);
	fmt.Println(arr2);

	// Slices

	// Array[startIndex(inclusive): endIndex(exclusive)]
	slices := arr2[2:5];
	fmt.Println(slices);

	slices2 := []int{}

	// Creating slices using make
	slice3 = make([]int, 5, 10);
	/*
		* In the above example complier will allocate 10 memory.
		* Here from index 0 to 4 compiler will be free to assign the default data
		* From 5 to 9 compiler will be create empty spaces. We need use append method add data.
		* EG [0,0,0,0,0,*,*,*,*]
			so length of above slice is 5 and capacity is 10.

	*/ 

	//  MAPs
	var stateCapital = make(map[string]string);
	// Map can also declared as
	// var mapExample1 map[key_dataType]value_dataType;
	stateCapital["Karnataka"] = "Bengaluru";
	stateCapital["Maharastra"] = "Mumbai";
	stateCapital["Kerala"] = "kerala";

	fmt.Println(stateCapital)

	for state := range stateCapital {
		fmt.Printf("Capital of", state, "is", stateCapital[state])
	} 

	capital, isCapital := stateCapital["Kerala"]

	if(isCapital) {
		fmt.Println("Capital is present and capital is", capital)
	} else {
		fmt.Println("Capital is not present ");
	}
}


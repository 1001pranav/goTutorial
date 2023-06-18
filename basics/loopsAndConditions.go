package main

import (
	"fmt"
	"strings"
)

func main() {
	var a, c int = 5, 1;

	if (a > c) {
		// F for format like %d 
		fmt.Printf("%d is greater than %d", a,c);
	} else if ( c > a) {
		fmt.Printf("%d is greater than %d", c, a);
	} else {
		fmt.Printf("%d is equal", a);
	}
	
	fmt.Printf("\n");
	var weeks = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"};

	var weekDay = "sunday";
	switch(strings.ToUpper(weekDay)) {
		case strings.ToUpper(weeks[0]):
			fmt.Println(weeks[0]);
			break;
		case strings.ToUpper(weeks[1]):
			fmt.Println(weeks[1]);
			break;
		case strings.ToUpper(weeks[2]):
			fmt.Println(weeks[2]);
			break;
		case strings.ToUpper(weeks[3]):
			fmt.Println(weeks[3]);
			break;
		case strings.ToUpper(weeks[4]):
			fmt.Println(weeks[4]);
			break;
		case strings.ToUpper(weeks[5]):
			fmt.Println(weeks[5]);
			break;
		case strings.ToUpper(weeks[6]):
			fmt.Println(weeks[6]);
			break;
		default:
			fmt.Println("Week Not found:");
	}
	
	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	
	for i := 0; i < len(numbers); i++ {
		if ( numbers[i] % 2 == 0 ) {
			fmt.Println("Even number", numbers[i]);
		} else {
			fmt.Println("Odd number", numbers[i]);
		}
	}

	var sum = 0;
	
	// _ is used to avoid not used warning
	for _, value := range numbers {
		sum = sum + value
		
	}

	fmt.Println("Sum is ", sum);
	const longestWord = "Pneumonoultramicroscopicsilicovolcanoconiosis";

	fmt.Println("Length of longest word is: ", len(longestWord));

	
}
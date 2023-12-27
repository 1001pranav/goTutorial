package main

import (
	"fmt"
	"strings"
)

func switchCondition(weekDay string) {
	var weeks = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"};

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
}

func invalidIfCondition() {
	var x int = 10;
	var y int = 20;
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
		fmt.Printf("%d is greater than %d\n", x, y);
	} else if x < y {
		fmt.Printf("%d is greater than %d\n", y, x);
	} else {
		fmt.Printf("%d is equal %d\n", x, y);
	}
}

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
	var weekDay = "sunday";
	switchCondition(weekDay);
	invalidIfCondition();
}
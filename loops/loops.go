package loop
import "fmt"

func loops() {
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

	index := 0;

	// Just specify the conditions
	for index < len(numbers) {
		fmt.Println("Numbers are: ", numbers[index]);
		index += 1;
	}
}
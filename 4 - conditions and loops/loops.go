package main
import (
	"fmt"
	"strings"
)

func basicForLoops(numbers []int) {
	for i := 0; i < len(numbers); i++ {
		if ( numbers[i] % 2 == 0 ) {
			fmt.Println("Even number", numbers[i]);
		} else {
			fmt.Println("Odd number", numbers[i]);
		}
	}
}

func rangeInForLoops(numbers []int) int {
	var sum int = 0;
	// _ is used to avoid not used warning
	for _, value := range numbers {
		sum += value

	}

	return sum;
}

func forWithOnlyCondition(numbers []int) map[string]int {
	var sumOddEven = map[string]int {
		"EVEN": 0,
		"ODD": 0,
	};

	var index int = 0;
	// Just specify the conditions
	for index < len(numbers) {
		if numbers[index] % 2 == 0 {
			sumOddEven["EVEN"] += numbers[index];
		} else {
			sumOddEven["ODD"] += numbers[index];
		}
		index += 1;
	}

	return sumOddEven;
}

func onlyForLoop(words string) map[string]int {
	// Using int because for Loop Will convert to ASCII
	wordsCount := make(map[string]int);

	var index int = 0;
	for {
		if (index>= len(words)) {
			break;
		}

		var charWord string = string(words[index]);
		charWord = 	strings.ToUpper(charWord);

		_, isValueExists := wordsCount[charWord];

		if isValueExists {
			wordsCount[charWord] += 1;
		} else {
			wordsCount[charWord] = 1;
		}

		index++;

	}

	return wordsCount;
}

func main() {
	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

	basicForLoops(numbers);

	var sum = rangeInForLoops(numbers);
	fmt.Println("Sum of numbers", sum);

	const LONGEST_WORD = "Pneumonoultramicroscopicsilicovolcanoconiosis";

	fmt.Println("Length of longest word is: ", len(LONGEST_WORD));
	var wordsCount = onlyForLoop(LONGEST_WORD);
	fmt.Println("Words count is", wordsCount);
	var sumOddEven = forWithOnlyCondition(numbers);
	fmt.Println("Sum of Even and Odd numbers are ", sumOddEven);

}
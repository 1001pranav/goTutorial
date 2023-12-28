package main

import (
	"fmt"
	"strings"
)

func main() {
	const LONGEST_WORD = "Pneumonoultramicroscopicsilicovolcanoconiosis";
	var wordsCount = make(map[string]int);

	for i := 0; i < len(LONGEST_WORD); i++ {
		/*
			LONGEST_WORD[i] -> Returns ASCII number in Bytes.
			string(LONGEST_WORD[i]) -> Converts to string from ASCII Bytes
		*/
		var character = string(LONGEST_WORD[i]);
		character = strings.ToUpper(character);

		_, isCharExists := wordsCount[character];

		if (isCharExists) {
			wordsCount[character] += 1;
		} else {
			wordsCount[character] = 1;
		}
	}
	fmt.Println(wordsCount);
}
package main

import (
	"fmt"
	e "exports"
)
/*
	- To working with the import export we need to  initialize the go module with following commands:
		`go mod init <projectName>`
	- When go module is initialized then we need to run go build command
		`go build`
	The above command will create an build and it will add the dependencies in `usr/local/go` folder.
*/
func main() {
	fmt.Println(e.PrimeNumbers(15));
}
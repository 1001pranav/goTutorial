package main
import "fmt"

/*
	1. ',' is not required for defining struct.
	2. we can create nested struct.
	syntax is
		type <structName> struct{
			<key/Field> <DataType>
			<key/Field> <DataType>
			<key/Field> <DataType>
			...
		}
*/
type person struct {
	firstName string
	age int
}

type families struct {
	father person
	mother person
	son person
}
func main() {
	fatherDetails := person {age:50, firstName: "Sundar"};
	motherDetails := person {firstName: "Sumithra", age: 40};
	sonDetails := person {firstName: "Suraj", age: 25};

	pavanFamily := families{father: fatherDetails, mother: motherDetails, son: sonDetails}
	fmt.Println(pavanFamily);
}
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

type employeeDetails struct {
	user person
	employeeID int

}

func (userDetails person) detail(){
	fmt.Println("First Name is :", userDetails.firstName);
	fmt.Println("Age is :", userDetails.age);
}

func (employee employeeDetails) detail()  {
	fmt.Println("Employee details are * * * * v");
	fmt.Println("Employee ID is :", employee.employeeID);
	employee.user.detail();

}

func main() {
	sundarDetails := person {age:50, firstName: "Sundar"};
	sundarDetails.detail();

	sunderJobDetails := employeeDetails{user: sundarDetails, employeeID: 1 };

	sunderJobDetails.detail();
	fmt.Println();
}
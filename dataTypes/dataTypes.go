package main
import "fmt";

func main() {

	var integer int = -10001;
	var floatVal float32 = 154.225;
	var message string = "String data type example";
	var isValid bool = true; // false;

	fmt.Println("string -", message, "\n", "boolean -", isValid);

	fmt.Println("AND - ",true && false);
	fmt.Println("OR - ", true || false);
	fmt.Println("NOT - ", !true);

	var sum, sub int;
	var div, mul, rem int;

	sum = integer + 2 ;
	sub = sum - 1;
	div = int(floatVal) / 2;
	mul = int(div) * sum;
	rem = mul % div;

	floatMul := floatVal/2;

	fmt.Println("Sum:", sum, "Multiplication:", mul, "division:", div, "reminder:", rem, "subtraction: ", sub);
	fmt.Println("Floating point Value for division", floatMul);
}
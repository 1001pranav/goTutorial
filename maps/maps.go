package main
import "fmt"

func main(){
	var countryCapital map[string]string = map[string]string {
		"India": "New Delhi",
		"Pakistan": "Islamabad",
		"China": "Beijing",
	}

	fmt.Println(countryCapital["India"]);

	// Deleing key in map
	delete(countryCapital, "Pakistan");

	countryCapital["Kailasa"] = "";
	// Accessing value with boolean
	fmt.Println(countryCapital);

	valKailasa, isKailasaExists := countryCapital["Kailasa"];
	fmt.Println("Is Kailasa Exists",isKailasaExists, "Value of Kailasa",  valKailasa);

	valEngland, isEnglandExists := countryCapital["England"];
	fmt.Println("England Exists", isEnglandExists, "Value of England is", valEngland);

	// Map can also declared as
	var stateCapital = make(map[string]string);

	// var mapExample1 map[key_dataType]value_dataType;
	stateCapital["Karnataka"] = "Bengaluru";
	stateCapital["Maharastra"] = "Mumbai";
	stateCapital["Kerala"] = "Thiruvananthapuram";

	fmt.Println(stateCapital)

	//We an only pass Key(state) here
	for state, city := range stateCapital {
		fmt.Println("Capital of", state, "is", city);
	}

	capital, isCapital := stateCapital["Kerala"]

	if(isCapital) {
		fmt.Println("Capital is present and capital is", capital)
	} else {
		fmt.Println("Capital is not present ");
	}
}
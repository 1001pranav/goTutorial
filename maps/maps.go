package maps
import "fmt"

func maps(){
	var stateCapital = make(map[string]string);
	// Map can also declared as
	// var mapExample1 map[key_dataType]value_dataType;
	stateCapital["Karnataka"] = "Bengaluru";
	stateCapital["Maharastra"] = "Mumbai";
	stateCapital["Kerala"] = "kerala";

	fmt.Println(stateCapital)

	for state := range stateCapital {
		fmt.Printf("Capital of", state, "is", stateCapital[state]);
	}

	capital, isCapital := stateCapital["Kerala"]

	if(isCapital) {
		fmt.Println("Capital is present and capital is", capital)
	} else {
		fmt.Println("Capital is not present ");
	}
}
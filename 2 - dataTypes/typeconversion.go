package main

import (
	"fmt"
	"math"
)

func main() {
	var i int = 42
	var f float64 = float64(i) * math.Pi
	var u uint = uint(f)
	fmt.Println(" Unsigned int ", u, "Float ", f, "Integer ", i)

	intNum := -450
	decimal := float64(intNum) * math.Pi
	unsignedInt := uint64(decimal)

	fmt.Println("Unsigned Integer ", unsignedInt, " Float value - ", decimal, " Integer Number", intNum)
}

package main

import (
	"fmt"
	"strings"
	"time"
)

var (
	FOOD_COST = map[string]int{
		"DRINKS":    25,
		"BREAKFAST": 75,
		"LUNCH":     125,
		"SNACKS":    60,
		"DINNER":    150,
	}
)

func switchCondition(weekDay string) {
	var weeks = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	switch strings.ToUpper(weekDay) {
	case strings.ToUpper(weeks[0]):
		fmt.Println(weeks[0])
		break
	case strings.ToUpper(weeks[1]):
		fmt.Println(weeks[1])
		break
	case strings.ToUpper(weeks[2]):
		fmt.Println(weeks[2])
		break
	case strings.ToUpper(weeks[3]):
		fmt.Println(weeks[3])
		break
	case strings.ToUpper(weeks[4]):
		fmt.Println(weeks[4])
		break
	case strings.ToUpper(weeks[5]):
		fmt.Println(weeks[5])
		break
	case strings.ToUpper(weeks[6]):
		fmt.Println(weeks[6])
		break
	default:
		fmt.Println("Week Not found:")
	}
}

func multipleConditions(foods []string) int {
	var totalSum int = 0
	for _, food := range foods {
		switch strings.ToUpper(food) {
		case "TEA", "COFFEE", "MILK":
			totalSum += FOOD_COST["DRINKS"]

		case "DOSA", "IDILY":
			totalSum += FOOD_COST["BREAKFAST"]

		case "FISH_MEALS", "VEG_MEALS":
			totalSum += FOOD_COST["LUNCH"]

		case "PANI PURI", "MASALA PURI", "SEV PURI":
			totalSum += FOOD_COST["SNACKS"]
		case "MEALS":
			totalSum += FOOD_COST["DINNER"]
		default:
			fmt.Println("Food not found")
		}
	}

	return totalSum
}

func switchTrue() {
	t := time.Now()
	//If any case becomes true then other cases won't be executing
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning")
	case t.Hour() < 15:
		fmt.Println("Good Afternoon")
	case t.Hour() < 18:
		fmt.Println("Good Evening")
	default:
		fmt.Println("Good Night")
	}
}

func main() {
	var orders = []string{"Tea", "Dosa"}
	var totalSum = multipleConditions(orders)
	fmt.Println("total bill amount: ", totalSum) // total bill amount: 100

	var weekDay = "sunday"
	switchCondition(weekDay)
	switchTrue()
}

package main

import "fmt"

var (
	mealCost float64
	tipPercent float64
	taxPercent float64
)

func main() {
	fmt.Scan(&mealCost)
	fmt.Scan(&tipPercent)
	fmt.Scan(&taxPercent)

	tip := mealCost*tipPercent/100
	tax := mealCost*taxPercent/100
	totalCost := mealCost + tip + tax
	fmt.Printf("The total meal cost is %v dollars.\n", int(totalCost+.5))
}
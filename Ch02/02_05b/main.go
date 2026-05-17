package main

import (
	"fmt"
	"math"
)

func main() {

	f1, f2, f3 := 23.5, 65.1, 76.3
	sum := f1 + f2 + f3
	fmt.Println("Float sum:", sum)

	sum = math.Round(sum*100)/100

	fmt.Printf("Rounded float sum: %v\n", sum)

	fmt.Println("Pie: ",math.Pi)
	var radius int = 5
	area := math.Pi * math.Pow(float64(radius), 2)
	fmt.Printf("Area of circle with radius %d: %.2f\n", radius, area)

}

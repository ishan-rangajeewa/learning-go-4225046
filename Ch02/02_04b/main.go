package main

import "fmt"

func main() {
	i1, i2, i3 := 12, 13, 14
	intsum := i1 + i2 + i3
	var f1, f2, f3 float64 = 12.5, 13.5, 14.5
	floatsum := f1 + f2 + f3

	answer := float64(i1) + f1
	fmt.Printf("Sum of f %v, %v and %v is: %v\n", f1, f2, f3, floatsum)
	fmt.Printf("Sum of i %v, %v and %v is: %v\n", i1, i2, i3, intsum)
	fmt.Printf("Sum of i+f %v and %v is: %v\n", i1, f1, answer)
	
}

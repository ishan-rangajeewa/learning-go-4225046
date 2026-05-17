package main

import (
	"fmt"
)

func main() {

	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	aNumber := 42

	fmt.Println(str1, str2, str3)
	stringLength, err := fmt.Println("Value is:", aNumber)
	if err == nil {
		fmt.Println("Length: ", stringLength)
	}
	fmt.Printf("Value is fPrint %v\n", aNumber)
	fmt.Printf("Value type : %T\n", aNumber)
}

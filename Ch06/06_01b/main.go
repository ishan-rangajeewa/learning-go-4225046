package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := "./test.txt"
	file, err := os.Create(fileName)
	checkError(err)
	defer file.Close()
	length, err := io.WriteString(file,"Hello World")
	checkError(err)
	fmt.Printf("Wrote %d bytes to file %s\n", length, fileName)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}
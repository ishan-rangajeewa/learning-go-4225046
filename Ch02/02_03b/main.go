package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number: ")
	str, _ := reader.ReadString('\n')
	println(str)
}

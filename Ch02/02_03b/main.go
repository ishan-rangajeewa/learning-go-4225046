package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number: ")
	str, _ := reader.ReadString('\n')
	println(str)

	fmt.Print("Enter number: ")
	str, _ = reader.ReadString('\n')
	aFloat,err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Printf("Value is: %v and type is: %T\n",aFloat,aFloat)
	}
}

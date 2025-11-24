package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// var s string
	// fmt.Scanln(&s)
	// fmt.Println(s)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter text:")
	str, _ := reader.ReadString('\n')
	fmt.Println(str)

	fmt.Print("enter a number")
	num, _ := reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(num), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("value of number: ", f)
	}

}

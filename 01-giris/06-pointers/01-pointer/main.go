package main

import "fmt"

func main() {
	//with default value
	var i int
	fmt.Println("address of i is: ", &i, "value at address is ", *&i)

	//with number value
	x := 23
	fmt.Println("address of x is: ", &x, "value at address is ", *&x)
}

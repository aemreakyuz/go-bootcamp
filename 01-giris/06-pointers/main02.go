package main

import (
	"fmt"
)

func main() {

	//pointer ı tanımlamak için * konulur
	var p *int

	if p != nil {
		fmt.Println("value of p is : ", *p)
	} else {
		fmt.Println("p is nil")
	}
	var v int = 42
	p = &v
	if p != nil {
		fmt.Println("value of p is : ", *p)
	} else {
		fmt.Println("p is nil")
	}
}

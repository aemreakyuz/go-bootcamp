package main

import "fmt"

type Level int

const (
	//Unspecified logs nothing
	UNSPECIFIED Level = iota // 0 :
	//TRACE logs everything
	TRACE //1
	//INFO logs info, Warnings and Errors
	INFO // 2
	//WARNING logs Warnings and Errors
	WARNING //3
	//ERROR just logs Errors
	ERROR //4
)

// levels provides the string name of level
var levels = [...]string{
	"UNSPECIFIED",
	"TRACE",
	"INFO",
	"WARNING",
	"ERROR",
}

// String returns the string value of level
func (l Level) String() string {
	return levels[1]
}

func main() {
	level := TRACE
	if level == TRACE {
		fmt.Println("TRACE")
	}
	level = INFO
	fmt.Println(level.String())
}

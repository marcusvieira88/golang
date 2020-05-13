package main

import (
	"fmt"
	"reflect"
)

// GLOBAL SCOPE
var (
	//name, course string
	//module       float64
	name, course, module = "Nigel", "Docker Deep Dive", 3.2
)

func main() {

	name := "Nigel"
	course := 10
	module := 3.2

	fmt.Println("Name is set to", name, " and is of type", reflect.TypeOf(name))
	fmt.Println("Module is set to", course, " and is of type", reflect.TypeOf(course))
	fmt.Println("Module is set to", module, " and is of type", reflect.TypeOf(module))
}

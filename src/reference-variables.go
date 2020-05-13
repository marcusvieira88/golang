package main

import (
	"fmt"
)

func main() {

	const env = "USERNAME"
	fmt.Println("Value of env variable is ", env)

	// := declare a new variable
	test := "Original Value"
	fmt.Println("Value of test variable is ", test)

	changeValue(&test)

	fmt.Println("Final value of test is ", test)
}

func changeValue(test *string) {

	// equal = changes only the value
	*test = "Changed Value"
	fmt.Println("Value Changed matehod is ", *test)
}

package main

import (
	"fmt"
)

func main() {

	testSlice := make([]string, 5, 10)
	fmt.Println(len(testSlice))
	fmt.Println(cap(testSlice))

	mySlice := []string{"a", "b", "c", "d", "e"}
	fmt.Println(mySlice)

	for index, value := range mySlice {
		fmt.Printf("index %d value %s \n", index, value)
	}

	newSlice := []string{"f", "g", "h"}
	mySlice = append(mySlice, newSlice...)
	fmt.Println(mySlice)
}

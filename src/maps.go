package main

import (
	"fmt"
)

func main() {

	testMap := make(map[string]int, 10)
	testMap["person1"] = 1
	testMap["person2"] = 2
	testMap["person3"] = 3

	for key, value := range testMap {
		fmt.Printf("key %s value %d \n", key, value)
	}

	testMap2 := map[string]int{"A": 10, "B": 20, "C": 30}
	for key, value := range testMap2 {
		fmt.Printf("key %s value %d \n", key, value)
	}
}

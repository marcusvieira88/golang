package main

import (
	"fmt"
)

func main() {

	firstValue := "11"
	secondValue := "23"

	if firstValue < secondValue {
		fmt.Println("secondValue value is higher than firstValue")
	} else if firstValue > secondValue {
		fmt.Println("firstValue value is higher than secondValue")
	} else {
		fmt.Println("Same value normal if")
	}

	if thirdValue, fourthValue := "40", "23"; thirdValue < fourthValue {
		fmt.Println("fourthValue value is higher than thirdValue")
	} else if thirdValue > fourthValue {
		fmt.Println("thirdValue value is higher than fourthValue")
	} else {
		fmt.Println("Same value if with initialization")
	}

	switch "marcus" {
	case "paulo":
		fmt.Println("Case Paulo")
	case "marcus":
		fmt.Println("Case Marcus")
	case "maria":
		fmt.Println("Case Maria")
	default:
		fmt.Println("Case Default")
	}
}

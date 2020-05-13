package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	players := []string{"Neymar", "Ronaldinho", "Rivaldo"}

	for index, value := range players {
		fmt.Println(value)
		fmt.Println(index)
	}

	for test := "10"; test == "10"; {
		fmt.Println(test)
		break
	}
}

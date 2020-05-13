package main

import (
	"fmt"
	"strings"
)

func main() {

	s1, s2 := "StringA", "StringB"

	fmt.Println(converter(s1, s2))

	test1, test2 := converter(s1, s2)

	fmt.Println(test1)
	fmt.Println(test2)
}

func converter(value1, value2 string) (str1, str2 string) {

	value1 = strings.Title(value1)
	value2 = strings.Title(value2)

	return value1, value2
}

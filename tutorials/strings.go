package main

import (
	"fmt"
	"strings"
)

func main() {

	str1 := "Hello"
	fmt.Printf("string1 %v:%T\n", str1, str1)
	fmt.Printf("string1 %v:%T\n", str1, str1)

	fmt.Println(strings.ToUpper(str1))
	fmt.Println(strings.Title(str1))

	lValue := "hello"
	uValue := "HELLO"

	fmt.Println("Equal? ", (lValue == uValue))
	fmt.Println("Equal-caseInsensitive? ", strings.EqualFold(lValue, uValue))

}

package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {

  // declare variable
  var s string
  fmt.Scanln(&s)
  fmt.Println(s)
  // if we type one two three
  // it will print only one because it breaks the space automatically

  reader := bufio.NewReader(os.Stdin)
  fmt.Println("Enter text: ")
  str, _ := reader.ReadString('\n')
  // _ is an error object _ means ignore the name
  // single quote means byte not string
  fmt.Println(str)

  fmt.Println("Enter Number: ")

  str, _ = reader.ReadString('\n')
  //dont need the : since the type has been declared
  f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
  //trim space removes \r\n
  //  64 = type of number we want to get back
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("Value of number : ", f)
  }


}

package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Hello from Go!")
	//P is uppercase because it's "exported"
	fmt.Println(strings.ToUpper("Hello from Go!"))
}

//   to run this because
//   go run hello-world.go

// to build use
// go build hello-world.go
// you will get an executable which you can use ./hello-world to run

// use gofmt to format the file
// gofmt -w hello-world.go

package main

import "fmt"

type MyInt int

func main() {
	var i interface{}
	var mine MyInt = 10
	i = mine
	fmt.Println(i.(MyInt) + 1)
}
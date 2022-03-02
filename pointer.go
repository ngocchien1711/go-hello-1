package main

import "fmt"

func main() {
	var x int
	pointerX := &x
	fmt.Println(pointerX)
	fmt.Println(*pointerX)

	var pointerY *int
	fmt.Println(pointerY)

	pointerZ := struct {
		name string
		age int
	} {
		name: "B",
		age: 20,
	}

	fmt.Println(pointerZ)
}

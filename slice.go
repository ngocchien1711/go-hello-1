package main

import "fmt"

type person struct {
	name string
	age  int
}

func addTo(base int, array ...int) []int {
	result := make([]int, 0, len(array))
	for _, v := range array {
		result = append(result, v+base)
	}
	return result
}
func main() {
	fmt.Println(addTo(10, 1, 2, 3))
}

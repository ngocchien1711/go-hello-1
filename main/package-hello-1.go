package main

import (
	"fmt"

	"gitlab.com/ngocchien/go-hello-1/formatter"
	"gitlab.com/ngocchien/go-hello-1/math"
)

func main() {
	num: math.Double(2)
	output := print.Format(num)
	fmt.Println(output)
}
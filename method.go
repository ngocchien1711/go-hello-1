package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: age %d", p.name, p.age)
}

func (p *Person) grow() {
	p.age++
}
func main() {
	p1 := Person{
		name: "A",
		age:  10,
	}

	p1.grow()

	pointer := &p1
	fmt.Println(pointer)
	fmt.Println(pointer.name)
	output := p1.String()
	fmt.Println(output)
}

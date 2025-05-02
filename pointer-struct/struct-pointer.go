package main

import (
	"fmt"
)

type Person struct {
	Name string
	age  int
}

//p := Person{}
//we can create like this as well

// take the pointer to the person and increase the age +1
func birthday(p *Person) {
	p.age++
}
func main() {

	p := Person{Name: "venthan", age: 35}
	fmt.Println("before birthday", p.age)
	birthday(&p)
	fmt.Println("after birthday", p.age)
}

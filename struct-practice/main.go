package main

import "fmt"

func main() {
	fmt.Println("Hello ,Go!")
	add()
	fmt.Println(sub())
	p := Person{Name: "john", Age: 35}
	fmt.Println(p.Name)

}

func add() {
	a := 3
	b := 4
	fmt.Println(a + b)
}

func sub() int {
	c := 34
	d := 0
	return c - d
}

type Person struct {
	Name string
	Age  int
}

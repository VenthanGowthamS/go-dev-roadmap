package main

import (
	"fmt"
)

func printShapeInfo(s shape) {
	fmt.Printf("Type: %T\n", s)
	fmt.Printf("area:%.2f\n", s.Area())
	fmt.Printf("perimeter:%.2f\n", s.Perimeter())
	fmt.Println("----------")

}

func main() {

	r := rectangle{width: 4, height: 6}
	s := square{side: 4}
	c := circle{radius: 6}

	shapes := []shape{r, s, c}

	for _, shape := range shapes {
		printShapeInfo(shape)
	}
}

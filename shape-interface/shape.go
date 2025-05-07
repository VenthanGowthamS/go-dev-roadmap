package main

import (
	"math"
)

//shape interface i want this behaviour have to be implemented whoever implements shape

type shape interface {
	Area() float64
	Perimeter() float64
}

//any type (struct) that claim to be shape must have two methods

//below are structs for rectangle

type rectangle struct {
	width, height float64
}

// implementing the interface

//for rectangle struct

func (r rectangle) Area() float64 {
	return r.width * r.height

}

func (r rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

//for cirlce
//struct for circle

type circle struct {
	radius float64
}

// implement share for circle
func (c circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// for square
type square struct {
	side float64
}

// implement the shape for square
func (s square) Area() float64 {
	return s.side * s.side
}

func (s square) Perimeter() float64 {
	return s.side * 4
}

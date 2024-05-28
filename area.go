package main

import "fmt"

type Shape interface {
	getArea() float64
}

type Square struct {
	sideLength float64
}

type Triangle struct {
	height float64
	base   float64
}

func (square Square) getArea() float64 {
	return square.sideLength * square.sideLength
}

func (triangle Triangle) getArea() float64 {
	return 0.5 * triangle.height * triangle.base
}

func printArea(shape Shape) {
	fmt.Println(shape.getArea())
}

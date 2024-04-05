package main

import "fmt"

type triangle struct {
	height float64
	base float64
}

type square struct {
	sideLength float64
}

type shape interface {
	printArea()
}

func main() {
	tr := triangle{height: 10, base: 10}
	tr.printArea()
	sq := square{sideLength: 10}
	sq.printArea()

}


func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) printArea() {
	fmt.Printf("Triangle area: %f\n",t.getArea())
}

func (s square) printArea() {
	fmt.Printf("Square area: %f\n",s.getArea())
}
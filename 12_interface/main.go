package main

import (
	"fmt" //Format Package
	"math"
)

// define interface
type Shape interface {
	area() float64
}
type Circle struct {
	x , y , radius float64
}
type Rectangle struct {
	x , y float64
}

func (c Circle) area() float64{
	return math.Pi * c.radius * c.radius
}
func (r Rectangle) area() float64 {
	return r.x * r.y
}
func  getArea(s Shape) float64 {
	return s.area()
}
func main() {
	circle := Circle { x:  0 , y : 0 , radius:  5}
	rectanlge := Rectangle{x : 10 , y : 5}
	// print new line
	fmt.Printf("Circle area is %f\n" , getArea(circle));
	fmt.Printf("Rectangle area is %f\n" , getArea(rectanlge));

}
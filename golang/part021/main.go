package main

import "fmt"

type Shape interface{
	area() float64
}

type MultiShape struct{
	shapes []Shape
}

func (m *MultiShape) area() float64{
	var area float64
	for _, s := range m.shapes{
		area += s.area()
	}
	return area
}

func totalArea(shapes ...Shape) float64{
	var area float64
	for _, s := range shapes{
		area += s.area()
	}
	return area
}

func main() {
	c := 10
	r := 20
	fmt.Println(totalArea(&c, &r))
}

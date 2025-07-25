package main

import (
	"chapter-7/interfaces"
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

type Person struct {
	Name string
}

type Android struct {
	Person //Android IS a Person indeed
	Model  string
}

// NewAndroid Constructor function for Android (Option 3)
func NewAndroid(name, model string) *Android {
	return &Android{
		Person: Person{Name: name},
		Model:  model,
	}
}

func (p Person) Talk() {
	fmt.Println("\nHi, my name is", p.Name)
}

func (a Android) Talk() {
	fmt.Println("\nHi, my name is", a.Name, "and my model is", a.Model)
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

func (rect Rectangle) Area() float64 {
	width := distance(rect.x1, rect.y1, rect.x2, rect.y1)
	height := distance(rect.x1, rect.y1, rect.x1, rect.y2)
	return width * height
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}

func main() {
	var rx1, ry1 float64 = 0, 0
	var rx2, ry2 float64 = 10, 10
	c := Circle{0, 0, 5}
	rect := Rectangle{rx1, ry1, rx2, ry2}
	fmt.Println("Rectangle area:", rect.Area())

	a := new(Person)
	a.Name = "Jam"
	a.Talk()

	jam := NewAndroid("Jam", "m3")
	jam.Talk()

	dummy := interfaces.TotalArea(&c, &rect)
	fmt.Println("\nTotal area:", dummy)

	multiShape := interfaces.MultiShapes{
		Shapes: []interfaces.Shape{
			Circle{0, 0, 2},
			Rectangle{0, 0, 10, 10},
		},
	}

	fmt.Println("\nTotal area of multi-shape:", interfaces.TotalArea(multiShape.Shapes...))
}

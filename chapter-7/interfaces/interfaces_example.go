package interfaces

type Shape interface {
	Area() float64
}

type MultiShapes struct {
	Shapes []Shape
}

func TotalArea(shapes ...Shape) (area float64) {
	for _, s := range shapes {
		area += s.Area()
	}
	return
}

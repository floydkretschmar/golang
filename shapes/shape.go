package shapes

import "fmt"

type Shape interface {
	GetName() string
	GetArea() float64
}

func PrintArea(s Shape) {
	fmt.Printf("The area of the %s is %f.\n", s.GetName(), s.GetArea())
}

func UseShapes() {
	triangle := Triangle{
		Name:   "My Triangle",
		Base:   5,
		Height: 4,
	}
	square := Square{
		Name:       "Your Square",
		SideLength: 3,
	}

	PrintArea(triangle)
	PrintArea(square)
}

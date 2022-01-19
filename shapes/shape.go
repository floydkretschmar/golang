package shapes

import "fmt"

type Shape interface {
	GetName() string
	GetArea() float64
}

func PrintArea(s Shape) {
	fmt.Printf("The area of the %s is %f.\n", s.GetName(), s.GetArea())
}

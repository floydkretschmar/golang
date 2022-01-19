package shapes

type Square struct {
	Name       string
	SideLength float64
}

func (s Square) GetArea() float64 {
	return s.SideLength * s.SideLength
}

func (s Square) GetName() string {
	return s.Name
}

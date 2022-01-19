package shapes

type Triangle struct {
	Name   string
	Base   float64
	Height float64
}

func (t Triangle) GetArea() float64 {
	return 0.5 * t.Base * t.Height
}

func (t Triangle) GetName() string {
	return t.Name
}

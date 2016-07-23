package interfaces

type Shape interface {
	Area() int
}

type Rectangle struct {
	length,breadth int
}

type Square struct {
	side int
}

type Hybrid struct {
	square Shape
	rectangle Shape
}

func (r Rectangle) Area() int {
	return r.length * r.breadth
}

func (square Square) Area() int {
	return square.side * square.side
}

func (h Hybrid) Area() int {
	return h.square.Area() + h.rectangle.Area()
}

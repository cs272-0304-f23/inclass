package main

type Shape interface {
	Area() int
	Perimeter() float32
}

type Rectangle {
	width int
	height int
}

func (r Rectangle) Area() int {
	return r.width * r.height
}

func main() {
	r := Rect{2, 3}
	a := r.Area()
}

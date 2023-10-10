package main

type Pet interface {
	Speak() string
	Toy() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "woof"
}

func (c Cat) Speak() {
	return "meow"
}

func Speaker(p Pet) {
	p.Speak()
}

func main() {
	dog := Dog{}
	Speaker(dog)
}

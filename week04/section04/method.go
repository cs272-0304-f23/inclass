package main

type Foo struct {
	i int
}

// PrintFoo is a method of the foo object
// f is a "receiver" in Go terminology, similar to self/this
func (f Foo) PrintFoo() {
	fmt.Println(f.i)
}

// Bar embeds a Foo, somewhat like inheriting a derived
// class from a base class in Java/C++
type Bar struct {
	Foo
	j int
}

func main() {
	var b Bar
	b.i = 0
}

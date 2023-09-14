package main

type Foo struct {
	i int
}

type Bar struct {
	Foo
	j int
}

func main() {
	var b Bar
	b.i = 0
}

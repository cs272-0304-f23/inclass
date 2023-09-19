package main

type Foo struct {
	i int
}

type Bar struct {
	Foo // embedded Foo in the Bar
	j int
}

func main() {
	var b Bar
	b.i = 0
}

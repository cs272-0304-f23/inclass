package main

import "testing"

func TestAdd2(t *testing.T) {
	a := 1
	b := 2
	want := a + b + 1
	got := Add2(a, b)
	if got != want {
		t.Errorf("Add2(%d, %d) got %d, want %d\n", a, b, got, want)
	}
}

func TestAdd3(t *testing.T) {
	got := Add3(1, 2, 3)
	want := 6
	if got != want {
		t.Errorf("Add3 got != want\n")
	}
}

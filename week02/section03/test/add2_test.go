package main

import "testing"

func TestAdd2(t *testing.T) {
	a := 2
	b := 3
	want := a + b
	got := Add2(a, b)
	if got != want {
		t.Errorf("Add2(%d, %d) = %d, want %d\n", a, b, got, want)
	}
}

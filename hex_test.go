package main

import "testing"

func TestHex(t *testing.T) {
	c := "#E4A75D"
	color, err := Hex(c, 1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(color)
}

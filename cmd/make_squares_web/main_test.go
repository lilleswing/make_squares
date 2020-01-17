package main

import (
	"testing"
)

func TestSquare(t *testing.T) {
	v1 := SquareIt(6)
	if v1 != "36" {
		t.Errorf("Wanted 36 got=%s", v1)
	}

}

package main

import "testing"

func TestBoat(t *testing.T) {
	result := Boat(2, 7)

	if result != 14 {
		t.Error()
	}

}

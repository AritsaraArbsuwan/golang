package main

import "testing"

func TestAnt(t *testing.T) {
	result := Ant(15, 7)

	if result != 8 {
		t.Error("Expected 8 , got", result)
	}

}

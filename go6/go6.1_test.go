package main

import "testing"

func TestDoremon(t *testing.T) {
	result := Doremon(5, 5)
	if result != 10 {
		t.Error("Expected 10 , got", result)
	}

}

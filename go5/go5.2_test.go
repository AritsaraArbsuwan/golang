package main

import "testing"

func TestArit(t *testing.T) {
	result := Arit(4, 2)
	if result != 2 {
		t.Error("Expected 2,got", result)
	}
}

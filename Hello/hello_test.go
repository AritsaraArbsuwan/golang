package main

import "testing"

func TestFind(t *testing.T) {
	t.Run("Hello World", func(t *testing.T) {
		want := true
		got := Find("Hello World")
		if got != want {
			t.Errorf("\ngot %v\n want %v", got, want)
		}
	})
	t.Run("Hello Solar System", func(t *testing.T) {
		want := false
		got := Find("Hello Soler System")
		if got != want {
			t.Errorf("\ngot %v\n want %v", got, want)
		}
	})
	t.Run("Hello Lunar", func(t *testing.T) {
		want := false
		got := Find("Hello Lunar")
		if got != want {
			t.Errorf("\ngot %v\n want %v", got, want)
		}
	})
}

package main

import "testing"

func TestBinary_search(t *testing.T) {
	t.Run("exist number", func(t *testing.T) {
		array := []int{1, 3, 5, 7, 9}
		got := Binary_search(array, 9)
		want := 4

		if got != want {
			t.Errorf("wanted %d, got %d", want, got)
		}
	})
	t.Run("not exist number", func(t *testing.T) {
		array := []int{1, 3, 5, 7, 9}
		got := Binary_search(array, 66)
		want := -1

		if got != want {
			t.Errorf("wanted %d, got %d", want, got)
		}
	})

}

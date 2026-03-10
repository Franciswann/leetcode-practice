package main

import "testing"

func TestBinary_search(t *testing.T) {
	array := []int{1, 3, 5, 7, 9}
	got := Binary_search(array, 3)
	want := 1

	if got != want {
		t.Errorf("wanted %d, got %d", want, got)
	}
}

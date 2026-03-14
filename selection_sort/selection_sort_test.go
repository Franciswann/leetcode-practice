package main

import "testing"

func TestSelectSort(t *testing.T) {
	arr := []int{100, 57, 61, 4, 7}
	got := SelectSort(arr)
	want := []int{4, 7, 57, 61, 100}

	if got != want {
		t.Errorf("wanted %v got %v", want, got)
	}
}

func TestFindSmallest(t *testing.T) {
	arr := []int{3, 4, 5, 1, 2}
	got := FindSmallest(arr)
	want := 1

	if got != want {
		t.Errorf("wanted %d got %d", want, got)
	}
}

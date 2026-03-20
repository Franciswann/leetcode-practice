package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	got := Sum([]int{2, 4, 6})
	want := 12
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCount(t *testing.T) {
	got := Count([]int{1, 2, 3, 4})
	want := 4

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestFind_Max(t *testing.T) {
	got := Find_Max([]int{2, 3, 100, 77})
	want := 100

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

}

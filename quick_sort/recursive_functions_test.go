package main

import (
	"reflect"
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

func TestQuick_Sort(t *testing.T) {
	got := Quick_Sort([]int{5, 4, 1, 3, 2})
	want := []int{1, 2, 3, 4, 5}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

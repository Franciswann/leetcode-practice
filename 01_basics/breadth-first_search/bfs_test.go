package main

import "testing"

// func TestB/

// func TestSearched_List(t *testing.T) {
// 	got := S
// }

func TestAlienVerifier(t *testing.T) {
	t.Run("Handle normal case", func(t *testing.T) {
		got := AlienVerifier("abcdefg")
		want := true
		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})

	// t.Run("Handle non-normal case", func(t *testing.T) {

	// })
}

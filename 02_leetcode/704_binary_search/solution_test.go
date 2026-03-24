package binary_search_704

import "testing"

func TestBinary_search(t *testing.T) {

	// test existing number and return its index
	t.Run("exist number", func(t *testing.T) {
		array := []int{1, 3, 5, 7, 9}
		got := Binary_search(array, 9)
		want := 4

		if got != want {
			t.Errorf("wanted %d, got %d", want, got)
		}
	})

	// test non-existing number and should return -1
	t.Run("not exist number", func(t *testing.T) {
		array := []int{1, 3, 5, 7, 9}
		got := Binary_search(array, 66)
		want := -1

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

}

package two_sum_001

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {

	t.Run("normal case test", func(t *testing.T) {
		got := twoSum([]int{1, 3, 2, 7}, 3)
		want := []int{0, 2}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("Exception Handling test", func(t *testing.T) {
		got := twoSum([]int{1, 2}, 5)
		if got != nil {
			t.Errorf("got %v, want nil", got)
		}
	})

}

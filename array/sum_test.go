package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Collection of 5 numbers", func(t *testing.T) {
		// this is array: size fixed
		// [5]int{...}

		// this is slice
		arr := []int{1, 2, 3, 4, 5}
		got := Sum(arr)
		want := 15

		if got != want {
			t.Errorf("expected %d but got %d", want, got)
		}
	})

	t.Run("Collection of any numbers", func(t *testing.T) {
		// this is array: size fixed
		// [5]int{...}

		// this is slice
		arr := []int{1, 2, 3}
		got := Sum(arr)
		want := 6

		if got != want {
			t.Errorf("expected %d but got %d", want, got)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	// slice can only be compared to nil, so this will be error
	// if got != want {
	// 	t.Errorf("expected %d but got %d", want, got)
	// }

	checkSum(t, got, want)

	// DeepEqual is not type safe, this will still be compiled
	// got := SumAll([]int{1, 2}, []int{0, 9})
	// want := "bob"
	// if !reflect.DeepEqual(got, want) {
	// 		t.Errorf("got %v want %v", got, want)
	// }
}

func TestSumAllTails(t *testing.T) {
	// Sum all Tails = take from 1 to the end
	t.Run("make the sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSum(t, got, want)
	})

	t.Run("make the sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSum(t, got, want)
	})
}

func checkSum(t testing.TB, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

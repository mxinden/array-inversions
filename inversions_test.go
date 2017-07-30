package inversions

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInversionsDefaultsToZero(t *testing.T) {
	if _, inv := Inversions([]int{}); inv != 0 {
		t.Error("expected 'Inversions' to return zero by default")
	}
}

var fileToInversionsTests = []struct {
	filePath           string
	expectedInversions int
}{
	{"examples/length-zero", 0},
	{"examples/length-one", 0},
	{"examples/length-two-no-inversions", 0},
	{"examples/length-two-one-inversion", 1},
	{"examples/length-three-one-inversion", 1},
}

func TestFileToInversions(t *testing.T) {
	for _, test := range fileToInversionsTests {
		amountInv, err := FileToInversions(test.filePath)
		if err != nil {
			t.Error(fmt.Sprintf("expected 'FileToInversions' with %v to not throw an error", test.filePath), err)
		}
		if amountInv != test.expectedInversions {
			t.Error(fmt.Sprintf(
				"expected 'FileToInversions with %v to return %v inversions, but returned %v instead",
				test.filePath,
				test.expectedInversions,
				amountInv,
			))
		}

	}
}

var countSplitInversionsTests = []struct {
	a          []int
	b          []int
	c          []int
	inversions int
}{
	// Without split inversion
	{[]int{}, []int{}, []int{}, 0},
	{[]int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}, 0},
	{[]int{1, 2}, []int{4, 5, 6}, []int{1, 2, 4, 5, 6}, 0},
	{[]int{1, 2, 3}, []int{5, 6}, []int{1, 2, 3, 5, 6}, 0},
	// With split inversion
	{[]int{2}, []int{1}, []int{1, 2}, 1},
	{[]int{1, 2, 4}, []int{3, 5, 6}, []int{1, 2, 3, 4, 5, 6}, 1},
	{[]int{4, 5, 6}, []int{1, 2, 3}, []int{1, 2, 3, 4, 5, 6}, 9},
}

func TestCountSplitInversions(t *testing.T) {
	for _, test := range countSplitInversionsTests {
		c, inversions := countSplitInversions(test.a, test.b)
		if reflect.DeepEqual(test.c, c) != true {
			t.Error(fmt.Sprintf("expected %v, got %v", test.c, c))
		}
		if inversions != test.inversions {
			t.Error(fmt.Sprintf("expected %v inversions, but got %v", test.inversions, inversions))
		}
	}
}

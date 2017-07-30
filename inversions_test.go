package inversions

import (
	"fmt"
	"testing"
)

func TestInversionsDefaultsToZero(t *testing.T) {
	t.Skip()
	if inv, err := Inversions([]string{}); err != nil && inv != 0 {
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

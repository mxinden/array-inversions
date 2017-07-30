package inversions

import (
	"testing"
)

func TestInversionsDefaultsToZero(t *testing.T) {
	if Inversions() != 0 {
		t.Error("expected 'Inversions' to return zero by default")
	}
}

func TestFileToInversionsWithEmptyFileReturnsZero(t *testing.T) {
	if FileToInversions("examples/length-zero.txt") != 0 {
		t.Error("expected 'FileToInversions' to return zero with empty file")
	}
}

func TestFileToInversionsWithLengthOneFileReturnsZero(t *testing.T) {
	if FileToInversions("examples/length-one.txt") != 0 {
		t.Error("expected 'FileToInversions' to return zero with length one file")
	}
}

package inversions

import (
	"testing"
)

func TestInversionsDefaultsToZero(t *testing.T) {
	if Inversions() != 0 {
		t.Error("expected 'Inversions' to return zero by default")
	}
}

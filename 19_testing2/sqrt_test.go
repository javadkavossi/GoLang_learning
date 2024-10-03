package sqrt

import (
	"testing"
)

func almostEqual(a, b float64) bool {
	return Abs(a-b) < 0.001
}

func TestSimple(t *testing.T) {
	val, err := Sqrt(2)
	if err != nil {
		t.Fatalf("Error in calculation: %v", err)
	}
	if !almostEqual(val, 1.414) {
		t.Fatalf("bad value - %f", val)
	}
}

func TestNegative(t *testing.T) {
	_, err := Sqrt(-2)
	if err == nil {
		t.Fatalf("expected error for negative input, got none")
	}
}

func BenchmarkSqrt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Sqrt(float64(i))
		if err != nil {
			b.Fatalf("Error in calculation: %v", err)
		}
	}
}

// Run (test benchMark) :  go test -v -bench .

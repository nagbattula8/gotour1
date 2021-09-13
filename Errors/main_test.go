package main

import (
	"testing"
)

func TestSqrt(t *testing.T) {
	tables := []struct {
		x float64
		n float64
	}{
		{4, 2},
		{9, 3},
		{-1, -1},
	}

	for _, table := range tables {
		total, err := Sqrt(table.x)

		if table.x < 0 && err == nil {
			t.Errorf("Error not generated")
		} else if table.x > 0 && total != table.n {
			t.Errorf("Sum of (%v) was incorrect, got: %v, want: %v.", table.x, total, table.n)
		}
	}
}

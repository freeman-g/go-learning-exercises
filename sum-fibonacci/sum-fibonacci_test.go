package main

import (
	"testing"
)

func TestSumFibs(t *testing.T) {
	// table driven test in the format of input:expected result
	tests := map[int]int{
		10:      10,
		1000:    1785,
		4000000: 4613732,
		4:       5,
		75024:   60696,
		75025:   135721,
	}

	for k, v := range tests {
		if sumFibs(k) != v {
			t.Errorf("sumFibs(%v) should return %v but returned %v", k, v, sumFibs(k))
		}
	}
}

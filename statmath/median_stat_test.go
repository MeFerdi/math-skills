package statmath

import (
	"math"
	"testing"
)

func TestMedian(t *testing.T) {
	tests := []struct {
		name   string
		data   []float64
		expect float64
	}{
		{"Single value", []float64{5}, 5},
		{"Even number of values", []float64{1, 2, 3, 4}, 3},
		{"Odd number of values", []float64{1, 2, 3, 4, 5}, 3},
		{"Negative values", []float64{-1, -2, -3, -4, -5}, -3},
		{"Mixed values", []float64{1, 2, 3, 4, 5, -1, -2, -3, -4, -5}, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Median(test.data)
			if math.Abs(actual-test.expect) > 0.00001 {
				t.Errorf("Median of %v is %f, but expected %f", test.data, actual, test.expect)
			}
		})
	}
}

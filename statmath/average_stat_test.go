package statmath

import (
	"math"
	"testing"
)

func TestAverage(t *testing.T) {
	tests := []struct {
		name   string
		data   []float64
		expect float64
	}{
		{"Single value", []float64{5}, 5},
		{"Multiple values", []float64{1, 2, 3, 4, 5}, 3},
		{"Negative values", []float64{-1, -2, -3, -4, -5}, -3},
		{"Mixed values", []float64{1, 2, 3, 4, 5, -1, -2, -3, -4, -5}, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Average(test.data)
			if math.Abs(actual-test.expect) > 0.00001 {
				t.Errorf("Average of %v is %f, but expected %f", test.data, actual, test.expect)
			}
		})
	}
}

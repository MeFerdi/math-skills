package statmath

import (
	"math"
	"testing"
)

func TestVariance(t *testing.T) {
	tests := []struct {
		name   string
		data   []float64
		expect float64
	}{
		{"Single value", []float64{5}, 0},
		{"Two values", []float64{1, 2}, 0.25},
		{"Three values", []float64{1, 2, 3}, 0.6666666666666666},
		{"Four values", []float64{1, 2, 3, 4}, 1.25},
		{"Five values", []float64{1, 2, 3, 4, 5}, 2},
		{"Ten values", []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 8.25},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Variance(test.data)
			if math.Abs(actual-test.expect) > 0.0001 {
				t.Errorf("Variance of %v is %f, but expected %f", test.data, actual, test.expect)
			}
		})
	}
}

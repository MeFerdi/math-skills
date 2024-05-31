package statmath

import (
	"math"
	"testing"
)

func TestStandardDeviation(t *testing.T) {
	tests := []struct {
		name   string
		data   []float64
		expect float64
	}{
		{"Single value", []float64{5}, 0},
		{"Two values", []float64{1, 2}, 0.5},
		{"Three values", []float64{1, 2, 3}, 0.816497},
		{"Four values", []float64{1, 2, 3, 4}, 1.118033988749895},
		{"Five values", []float64{1, 2, 3, 4, 5}, 1.4142135623730951},
		{"Ten values", []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2.872281},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := StandardDeviation(test.data)
			if math.Abs(actual-test.expect) > 0.00001 {
				t.Errorf("Standard Deviation of %v is %f, but expected %f", test.data, actual, test.expect)
			}
		})
	}
}

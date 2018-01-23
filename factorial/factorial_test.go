package factorial

import (
	"testing"
)

var tests = []struct {
	value  int
	result int
}{
	{0, 0},
	{1, 1},
	{2, 2},
	{3, 6},
}

func TestCalculate(t *testing.T) {
	for _, test := range tests {
		expected := test.result
		actual := Calculate(test.value)
		if actual != expected {
			t.Errorf("Factorial(%d): expected %d, actual %d", test.value, test.result, actual)
		}
	}
}

package string2color

import (
	"testing"
)

func Test_Interpolate(t *testing.T) {
	tcs := []struct {
		name   string
		in     uint
		inMin  uint
		inMax  uint
		out    uint
		outMin uint
	}{
		{"midpoint", 5, 50, 0, 10, 0, 100},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := Interpolate(tc.in, 0, 10, 0, 100)
			if actual != tc.out {
				t.Logf("expected %d, got %d", tc.out, actual)
				t.Fail()
			}
		})
	}
}

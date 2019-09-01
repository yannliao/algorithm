package intro_6006

import "testing"

func TestPeakFinding2D(t *testing.T) {
	var tests = []struct {
		input [][]int
		want  []int
	}{
		{
			input: [][]int{{10, 14, 15, 16}, {8, 13, 9, 17}, {10, 12, 11, 19}, {10, 11, 21, 20}},
			want:  []int{3, 2},
		},
	}
	for _, test := range tests {
		if x, y := PeakFinding2D(test.input, 0, len(test.input)-1); x != test.want[0] || y != test.want[1] {
			t.Errorf("PeakFinding2D(%v,0,%d) = %d, %d", test.input, len(test.input)-1, x, y)
		}
	}
}

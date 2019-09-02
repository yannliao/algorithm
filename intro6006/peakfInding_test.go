package intro6006

import "testing"

func TestPeakFinding(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{
			input: []int{6, 7, 4, 3, 2, 1, 4, 5},
			want:  1,
		},
		{
			input: []int{1, 2, 3, 4, 5, 6},
			want:  5,
		},
		{
			input: []int{1, 2, 3, 1},
			want:  2,
		},
	}
	for _, test := range tests {
		if got := PeakFinding(test.input, 0, len(test.input)-1); got != test.want {
			t.Errorf("PeakFinding(%v,0,%d) = %d", test.input, len(test.input)-1, got)
		}
	}
}

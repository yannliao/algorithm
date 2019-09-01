package sorting_6006

import "testing"

func TestMergeSort(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{
			input: []int{8, 2, 4, 9, 3, 6},
			want:  []int{2, 3, 4, 6, 8, 9},
		},
		{
			input: []int{3, 5, 1, 2, 4, 7, 6},
			want:  []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			input: []int{2, 2},
			want:  []int{2, 2},
		},
		{
			input: []int{3},
			want:  []int{3},
		},
		{
			input: []int{},
			want:  []int{},
		},
	}
	for _, test := range tests {
		if got := MergeSort(test.input); !equal(got, test.want) {
			t.Errorf("MergeSort(%v)= %v", test.input, got)
		}
	}
}

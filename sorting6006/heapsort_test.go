package sorting6006

import "testing"

func TestHeapSort(t *testing.T) {
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
			input: []int{3},
			want:  []int{3},
		},
		{
			input: []int{},
			want:  []int{},
		},
	}
	for _, test := range tests {
		t.Logf("HeapSort, sorting %v", test.input)
		if HeapSort(&test.input); !equal(test.input, test.want) {
			t.Errorf("HeapSort Fail result = %v", test.input)
		}
	}
}

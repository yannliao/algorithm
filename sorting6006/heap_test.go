package sorting6006

import "testing"

func TestHeap(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{
			input: []int{8, 2, 4, 9, 3, 6},
			want:  []int{9, 8, 6, 2, 3, 4},
		},
	}
	for _, test := range tests {
		t.Logf("Test Heap Insert %v", test.input)
		var h MaxHeap
		for _, v := range test.input {
			h.Insert(v)
		}
		if !equal(h, test.want) {
			t.Errorf("Test Heap Insert Fail result = %v, want = %v", h, test.want)
		}
	}
}
func TestExtractMax(t *testing.T) {
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
		var h MaxHeap

		for _, v := range test.input {
			h.Insert(v)
		}
		got := h
		for len(h) > 0 {
			h.ExtractMax()
		}
		if !equal(got, test.want) {
			t.Errorf("Test Heap ExtractMax Fail result = %v, want = %v", got, test.want)
		}
	}
}

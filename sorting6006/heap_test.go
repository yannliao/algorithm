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
			want:  []int{8, 4, 6, 2, 3},
		},
	}
	for _, test := range tests {
		t.Logf("Test Heap ExtractMax %v", test.input)
		var h MaxHeap
		for _, v := range test.input {
			h.Insert(v)
		}
		got := h.ExtractMax()
		if got != 9 {
			t.Errorf("Test Heap ExtractMax Fail got = %v, want = %v", got, 9)
		}
		if !equal(h, test.want) {
			t.Errorf("Test Heap ExtractMax Fail result = %v, want = %v", h, test.want)
		}
	}

}

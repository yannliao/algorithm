package sorting_6006

import "testing"

func TestInsertSort(t *testing.T) {
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
		t.Logf("InsertSort, sorting %v", test.input)
		if InsertSort(&test.input); !equal(test.input, test.want) {
			t.Errorf("InsertSort Fail result = %v", test.input)
		}
	}
}

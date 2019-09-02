package sorting_6006

func parent(i int) int {
	return (i - 1) / 2
}
func left(i int) int {
	return 2*i + 1
}
func right(i int) int {
	return 2 * (i + 1)
}

func buidmaxheap(s *[]int) {
	i := len(*s)/2 - 1
	for i >= 0 {
		maxheapify(s, len(*s), i)
		i--
	}
}
func maxheapify(s *[]int, size int, i int) {
	// p := parent(i)
	l := left(i)
	r := right(i)
	// length := len(*s)
	max := i
	if l <= size && max < (*s)[l] {
		max = l
	}
	if r <= size && max < (*s)[r] {
		max = l
	}
	if max != i {
		(*s)[i], (*s)[max] = (*s)[max], (*s)[i]
		maxheapify(s, size, max)
	}
}

func HeapSort(s []int) {
	size := len(s)
	buidmaxheap(&s)
	for size > 0 {
		s[0], s[size-1] = s[size-1], s[0]
		size--
		maxheapify(&s, size, 0)

	}
}

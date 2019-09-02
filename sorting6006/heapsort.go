package sorting6006

func parent(i int) int {
	return (i - 1) / 2
}
func left(i int) int {
	return 2*i + 1
}
func right(i int) int {
	return 2 * (i + 1)
}

func maxheapify(s *[]int, size int, i int) {
	// p := parent(i)
	l := left(i)
	r := right(i)
	// length := len(*s)
	max := i
	if l < size && (*s)[max] < (*s)[l] {
		max = l
	}
	if r < size && (*s)[max] < (*s)[r] {
		max = r
	}
	if max != i {
		(*s)[i], (*s)[max] = (*s)[max], (*s)[i]
		maxheapify(s, size, max)
	}
}

func buidmaxheap(s *[]int) {
	size := len(*s)
	i := len(*s)/2 - 1
	for i > -1 {
		maxheapify(s, size, i)
		i--
	}
}

//HeapSort algorithm
func HeapSort(s *[]int) {
	buidmaxheap(s)
	size := len(*s)
	for i := size - 1; i > 0; i-- {
		(*s)[0], (*s)[i] = (*s)[i], (*s)[0]
		size--
		maxheapify(s, size, 0)
	}
}

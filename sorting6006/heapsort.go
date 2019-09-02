package sorting6006

func maxheapify(p *[]int, size int, i int) {
	s := *p
	l := left(i)
	r := right(i)
	max := i
	if l < size && s[max] < s[l] {
		max = l
	}
	if r < size && s[max] < s[r] {
		max = r
	}
	if max != i {
		s[i], s[max] = s[max], s[i]
		maxheapify(p, size, max)
	}
}

func buidmaxheap(p *[]int) {
	s := *p
	size := len(s)
	i := len(s)/2 - 1
	for i >= 0 {
		maxheapify(p, size, i)
		i--
	}
}

//HeapSort algorithm
func HeapSort(p *[]int) {
	s := *p
	buidmaxheap(p)
	size := len(s)
	for i := size - 1; i > 0; i-- {
		(s)[0], (s)[i] = (s)[i], (s)[0]
		size--
		maxheapify(p, size, 0)
	}
}

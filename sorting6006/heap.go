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

//MaxHeap data structure
type MaxHeap []int

func (p *MaxHeap) maxheapify(i int) {
	l := left(i)
	r := right(i)
	size := len(*p)
	max := i
	if l < size && (*p)[max] < (*p)[l] {
		max = l
	}
	if r < size && (*p)[max] < (*p)[r] {
		max = r
	}
	if max != i {
		(*p)[max], (*p)[i] = (*p)[i], (*p)[max]
		p.maxheapify(max)
	}
}

//Insert any key to max heap
func (p *MaxHeap) Insert(k int) {
	i := len(*p)
	f := parent(i)
	*p = append(*p, k)
	for i > 0 && (*p)[i] > (*p)[f] {
		(*p)[i], (*p)[f] = (*p)[f], (*p)[i]
		i = f
		f = parent(i)
	}
}

//ExtractMax of MaxHeap
func (p *MaxHeap) ExtractMax() int {
	l := len(*p)
	v := (*p)[0]
	(*p)[0] = (*p)[l-1]
	*p = (*p)[:l-1]
	p.maxheapify(0)
	return v
}

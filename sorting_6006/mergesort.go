package sorting_6006

//MergeSort algorithm
func MergeSort(s []int) []int {
	if len(s) <= 1 {
		return s
	}
	mid := len(s) / 2
	left := s[:mid]
	right := s[mid:]
	re := merge(MergeSort(left), MergeSort(right))
	return re
}

//merge two sorted array
func merge(a []int, b []int) []int {
	i := 0
	j := 0
	var re []int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			re = append(re, a[i])
			i++
		} else {
			re = append(re, b[j])
			j++
		}
	}
	if i < len(a) {
		re = append(re, a[i:]...)
	}
	if j < len(b) {
		re = append(re, b[j:]...)
	}
	return re
}

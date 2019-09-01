package sorting_6006

func InsertSort(s *[]int) {
	if len(*s) <= 1 {
		return
	}
	for j := 1; j < len(*s); j++ {
		tmp := (*s)[j]
		i := j - 1
		for i >= 0 && (*s)[i] > tmp {
			(*s)[i+1] = (*s)[i]
			i--
		}
		(*s)[i+1] = tmp
	}
	return
}

func equal(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

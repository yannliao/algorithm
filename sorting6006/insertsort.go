//Package sorting6006 implements sorting algorithms in sort & trees section of MIT 6.006.
package sorting6006

//InsertSort algorithm
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

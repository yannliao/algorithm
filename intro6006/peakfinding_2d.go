package intro6006

//PeakFinding2D algorithm
func PeakFinding2D(s [][]int, low int, high int) (int, int) {
	mid := (low + high) / 2
	max := findMax(s[mid])
	if mid > 0 && s[mid][max] < s[mid-1][max] {
		return PeakFinding2D(s, low, mid-1)
	} else if mid < len(s)-1 && s[mid][max] < s[mid+1][max] {
		return PeakFinding2D(s, mid+1, high)
	}
	return mid, max
}
func findMax(s []int) int {
	j := 0
	for i, value := range s {
		if value > s[j] {
			j = i
		}
	}
	return j
}

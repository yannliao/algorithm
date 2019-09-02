//Package intro6006 implements PeakFinding and PeakFinding2D in introduction section of MIT 6.006.
package intro6006

//PeakFinding algorithm
func PeakFinding(s []int, low int, high int) int {
	mid := (low + high) / 2
	if mid > 0 && s[mid] < s[mid-1] {
		return PeakFinding(s, low, mid-1)
	} else if mid < len(s)-1 && s[mid] < s[mid+1] {
		return PeakFinding(s, mid+1, high)
	}
	return mid
}

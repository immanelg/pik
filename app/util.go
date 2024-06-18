package app

func clamp(value int, left int, right int) int {
	return max(left, min(value, right))
}

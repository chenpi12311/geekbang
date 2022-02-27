package plus_one

func plusOne(digits []int) []int {
	c := 1
	for i := len(digits) - 1; i >= 0 && c != 0; i-- {
		num := (digits[i] + c) % 10
		if num == 0 {
			c = 1
		} else {
			c = 0
		}
		digits[i] = num
	}
	if c == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
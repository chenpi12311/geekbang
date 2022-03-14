package week2

// 560. 和为 K 的子数组
func subarraySum(nums []int, k int) int {
	s := make([]int, len(nums)+1, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		s[i] = s[i-1] + nums[i-1]
	}
	var count int
	for i := 0; i < len(s); i++ {
		for j := i+1; j < len(s); j++ {
			if s[j] - s[i] == k {
				count++
			}
		}
	}
	return count
}

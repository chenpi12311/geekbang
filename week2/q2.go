package week2

// 数组的度
func findShortestSubArray(nums []int) int {
	m := make(map[int][]int, 0)
	for index, num := range nums {
		if arr, ok := m[num]; ok {
			arr = append(arr, index)
			m[num] = arr
		} else {
			arr = make([]int, 0)
			arr = append(arr, index)
			m[num] = arr
		}
	}
	// fmt.Println(m)
	max := -1
	minLen := 50001
	for _, v := range m {
		if len(v) >= max {
			start, end := v[0], v[len(v)-1]
			if len(v) > max {
				minLen = end - start + 1
			} else {
				if end-start+1 < minLen {
					minLen = end - start + 1
				}
			}
			max = len(v)
		}
	}
	return minLen
}
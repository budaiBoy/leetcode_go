package leetcode

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	m, res := make(map[int]bool), 0
	for _, v := range nums {
		m[v] = true
	}
	for k := range m {
		if !m[k-1] && !m[k+1] {
			delete(m, k)
		}
	}
	if len(m) == 0 {
		return 1
	}

	for k := range m {
		if m[k-1] {
			continue
		}
		temp := 1
		for i := k + 1; m[i]; i++ {
			temp++
		}
		res = max(res, temp)
	}

	return res
}

func longestConsecutive2(nums []int) int {
	m, res := make(map[int]int), 0
	for _, v := range nums {
		if _, ok := m[v]; ok {
			continue
		}
		left, right, sum := 0, 0, 0
		if val, ok := m[v-1]; ok {
			left = val
		}
		if val, ok := m[v+1]; ok {
			right = val
		}
		sum = left + right + 1
		m[v] = sum
		res = max(res, sum)
		m[v-left] = sum
		m[v+right] = sum
	}
	return res

}

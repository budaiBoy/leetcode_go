package leetcode

func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	i := 0
	j := 0
	for j < len(nums) {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		j++
	}
}

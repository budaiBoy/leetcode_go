package leetcode

func maxArea(height []int) int {
	maxArea, left, right := 0, 0, len(height)-1
	for left < right {
		weight := right - left
		area := weight * min(height[left], height[right])

		if area > maxArea {
			maxArea = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

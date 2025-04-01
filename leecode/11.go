package main

func maxArea(height []int) int {

	var maxArea int
	var left int
	for i := range height {
		if maxArea > height[i] {
			left = left
			maxArea = maxArea
		} else {
			left = i
			maxArea = height[i]
		}
	}
}

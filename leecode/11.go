package main

import "fmt"

func maxArea(height []int) int {
	var maxArea int
	n := len(height)

	for i := 0; i < n; i++ {
		for j := 1; j < n; j++ {
			area := height[i] * height[j] * (i - j)
			maxArea = max(area, maxArea)
		}
	}

	fmt.Println(maxArea)
	return maxArea
}

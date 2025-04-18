package main

import "fmt"

func main() {
	trap([]int{4, 2, 0, 3, 2, 5})
}

func trap(height []int) int {

	sumWhite := 0
	max1 := 0
	for i := 0; i < len(height); i++ {
		max1 = max(max1, height[i])
		sumWhite += height[i]
	}

	n := len(height)
	if n == 0 {
		return 0
	}

	h := 1

	sumAll := 0
	left := 0
	right := n - 1

	for h <= max1 {

		for height[left] < h {
			left++
		}
		for height[right] < h {
			right--
		}
		if left == right {
			sumAll += 1
		} else {
			sumAll += right - left + 1
		}

		fmt.Println(sumAll, h, left, right)
		h++
	}

	return sumAll - sumWhite
}

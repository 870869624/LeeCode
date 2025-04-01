package main

func threeSum(nums []int) [][]int {

	result := make([][]int, 0)
	sum := make([]int, 0)

	n := len(nums)

	if n < 3 {
		return result
	} else if n == 3 {
		if nums[0]+nums[1]+nums[2] == 0 {
			result = append(result, []int{nums[0], nums[1], nums[2]})
		}
		return result
	}

	return result
}

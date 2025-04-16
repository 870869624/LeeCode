package main

func subarraySum(nums []int, k int) int {
	preSum := make(map[int]int)
	preSum[0] = 1

	var sum int
	var count int
	for _, v := range nums {
		sum += v
		if v, ok := preSum[sum-k]; ok {
			count += v
		}
		preSum[sum]++
	}
	return count
}

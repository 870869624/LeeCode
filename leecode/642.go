package main

import "fmt"

func main() {
	nums := []int{8860, -853, 6534, 4477, -4589, 8646, -6155, -5577, -1656, -5779, -2619, -8604, -1358, -8009, 4983, 7063, 3104, -1560, 4080, 2763, 5616, -2375, 2848, 1394, -7173, -5225, -8244, -809, 8025, -4072, -4391, -9579, 1407, 6700, 2421, -6685, 5481, -1732, -8892, -6645, 3077, 3287, -4149, 8701, -4393, -9070, -1777, 2237, -3253, -506, -4931, -7366, -8132, 5406, -6300, -275, -1908, 67, 3569, 1433, -7262, -437, 8303, 4498, -379, 3054, -6285, 4203, 6908, 4433, 3077, 2288, 9733, -8067, 3007, 9725, 9669, 1362, -2561, -4225, 5442, -9006, -429, 160, -9234, -4444, 3586, -5711, -9506, -79, -4418, -4348, -5891}
	res := findMaxAverage(nums, 93)
	fmt.Println(res)
}
func findMaxAverage(nums []int, k int) float64 {
	var sum int
	var ave float64
	var maxAve float64
	var flag int
	for i, v := range nums {
		sum += v
		if i < k-1 {
			continue
		}
		flag++
		ave = float64(sum) / float64(k)
		maxAve = max(ave, maxAve)

		if flag == 1 {
			maxAve = ave
		}

		sum -= nums[i-k+1]
	}
	return maxAve
}

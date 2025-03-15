package main

func getAverages(nums []int, k int) []int {

	var sum int64
	var ave int64

	var res []int
	var res1 []int
	for i, v := range nums {
		sum += int64(v)

		if i < k {
			res = append(res, -1)
			continue
		}

		if i < 2*k {
			continue
		} else if i+k > len(nums)-1 {
			res1 = append(res1, -1)
		}

		ave = sum / (2*int64(k) + 1)

		res = append(res, int(ave))

		sum -= int64(nums[i-2*k])
	}
	res = append(res, res1...)
	return res
}

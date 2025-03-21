package main

import "fmt"

func main() {
	maxFreeTime(10, 2, []int{0, 2, 9}, []int{1, 4, 10})
}
func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) int {
	var ans int
	n := len(startTime) //总共有几段会议

	get := func(i int) int {
		if i == 0 {
			return startTime[0] //第一段会议前的空余时间
		}

		if i == n {
			return eventTime - endTime[n-1] //中间可获得的空余时间
		}
		return startTime[i] - endTime[i-1] //最后一段会议到时间结束的空余时间
	}

	s := 0

	//三段会议，采用n+1便利可以获得0-3,意思就是三段会议中间有四段空隙
	for i := range n + 1 {
		s += get(i)
		if i < k {
			continue
		}

		ans = max(ans, s)
		s -= get(i - k)
		fmt.Println("get", get(i), -get(i-k), ans, i)
	}

	fmt.Println("max", ans)
	return ans

}

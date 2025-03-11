package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	// info := make(chan time.Time, 10)
	var input string
	const layout = "Jan 2, 2006 at 3:04pm (MST)"
	t := time.Date(2009, time.November, 10, 15, 0, 0, 0, time.Local)
	fmt.Println(t.Format(layout))
	fmt.Println(t.UTC().Format(layout))
	fmt.Println("输入")
	fmt.Scanln(&input)
	second, h, m, s, err := inputString(input)
	if err != nil {
		fmt.Println(err)
	}
	// for {
	// 	go add(second, h, m, s, info)
	// 	select {
	// 	case <-info:
	// 		ll := <-info
	// 		fmt.Println(ll, "1111111111111")
	// 	default:
	// 		fmt.Println("1111")
	// 	}
	// }

	for i := second; i > 0; i-- {
		time.Sleep(time.Second)
		s = s - 1
		if s == -1 {
			m = m - 1
			s = 59
			if m == -1 {
				h = h - 1
				m = 59
				if h == -1 {
					return
				}
			}
		}
		if h < 10 {
			if m < 10 {
				if s < 10 {
					fmt.Printf("还剩0%d时0%d分0%d秒\n", h, m, s)
				} else {
					fmt.Printf("还剩0%d时0%d分%d秒\n", h, m, s)
				}
			} else if s < 10 {
				fmt.Printf("还剩0%d时%d分0%d秒\n", h, m, s)
			} else {
				fmt.Printf("还剩0%d时%d分%d秒\n", h, m, s)
			}

		} else if h > 10 {
			if m < 10 {
				if s < 10 {
					fmt.Printf("还剩0%d时0%d分0%d秒\n", h, m, s)
				} else {
					fmt.Printf("还剩0%d时0%d分%d秒\n", h, m, s)
				}
			} else if s < 10 {
				fmt.Printf("还剩0%d时%d分0%d秒\n", h, m, s)
			} else {
				fmt.Printf("还剩%d时%d分%d秒\n", h, m, s)
			}
		}
	}

	timer("00:01:30")
}

// func add(second int, h int, m int, s int, ch chan time.Time) chan string {
// 	for i := second; i > 0; i-- {
// 		time.Sleep(time.Second)
// 		s = s - 1
// 		if s == -1 {
// 			m = m - 1
// 			s = 59
// 			if m == -1 {
// 				h = h - 1
// 				m = 59
// 				if h == -1 {
// 					return nil
// 				}
// 			}
// 		}
// 		time := fmt.Sprintf("还剩%d时%d分%d秒", h, m, s)
// 		ch <- time
// 	}
// 	return ch
// }

// 先输入事件再计算有多少时间
func inputString(str string) (int, int, int, int, error) {
	a := make([]int, 0)
	num := strings.Split(str, ":")

	for i := 0; i < len(num); i++ {
		c, err := strconv.Atoi(num[i])
		if err != nil {
			return 0, 0, 0, 0, err
		}
		a = append(a, c)
	}

	second := 3600*a[0] + 60*a[1] + a[2]
	return second, a[0], a[1], a[2], nil
	// nums := make([]int, 0)
	// num := strings.Split(str, "")
	// for i := 0; i < len(num); i++ {
	// 	a, err := strconv.Atoi(num[i])
	// 	if err != nil {
	// 		return time.Now(), err
	// 	}
	// 	nums = append(nums, a)
	// }

	// length := len(nums)
	// switch length {
	// case 4:
	// 	a := nums[length-4]*1000 + nums[length-3]*100 + nums[length-2]*10 + nums[length-1]
	// 	fmt.Println("aaaaaaaaaaaa", a)
	// case 3:
	// case 2:
	// case 1:
	// }

	// fmt.Println(nums)
	// return time.Now(), nil
}

var base, _ = time.Parse("15:04:05", "00:00:00")

func timer(str string) {
	t := Parase(str)
	fmt.Println(t)
	for {
		if t < 0 {
			return
		}
		time.Sleep(time.Second)
		t = t - time.Second
		fmt.Println(Format(t))
	}
}

func Format(t time.Duration) string {
	arg := base.Add(t)
	str := arg.Format("15:04:05")
	return str
}

func Parase(input string) time.Duration {
	t, _ := time.Parse("15:04:05", input)
	t1 := t.Sub(base)
	return t1
}

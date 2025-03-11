package duilie

import "fmt"

type shulie struct {
	max   int
	arr   [5]int
	front int
	fear  int
}

// 完成一个简单的入站出站操作
func (s *shulie) add(v int) {
	if s.fear == s.max-1 {
		fmt.Println("已经满了")
		return
	}
	s.fear++
	s.arr[s.fear] = v
}
func (s *shulie) Show() {
	for i := s.fear; i > s.front; i-- {
		fmt.Println(s.arr[i])
	}
}
func (s *shulie) delete(v int) {
	if s.fear == s.front+1 {
		fmt.Println("已经空了")
		return
	}
	for k, v1 := range s.arr {
		if v1 == v {
			s.arr[k] = 0
		}
		fmt.Println("没有这个数")
		return
	}
}
func main() {
	var a int
	b := &shulie{
		max:   5,
		front: -1,
		fear:  -1,
	}
	for {

		fmt.Println("\n1.增加栈")
		fmt.Println("2.减少栈")
		fmt.Println("3.展示")
		fmt.Scanln(&a)
		fmt.Printf("\n")
		switch a {
		case 1:
			fmt.Println("请输入入站的数")
			fmt.Scanln(&a)
			b.add(a)
		case 2:
			fmt.Println("请输入删除的数")
			fmt.Scanln(&a)
			b.delete(a)
		case 3:
			b.Show()
		}
	}
}

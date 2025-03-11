package main

import (
	"fmt"
	"math/rand"
)

type sort func([]int) []int

func paixu(num []int, sort sort) []int {
	return sort(num)
}

func maopaoSort(num []int) []int {
	for i := 0; i < len(num); i++ {
		for j := 0; j < len(num)-1; j++ {
			if num[j] > num[j+1] {
				num[j], num[j+1] = num[j+1], num[j]
			}
		}
	}
	return num
}
func xuanzeSort(num []int) []int {
	for i := 0; i < len(num); i++ {
		for j := i + 1; j < len(num); j++ {
			if num[i] > num[j] {
				min := num[i]
				num[i] = num[j]
				num[j] = min
			}
		}
	}
	return num
}
func main() {
	s := make([]int, 0)
	for i := 0; i < 100; i++ {
		s = append(s, rand.Intn(100))
	}
	num1 := []int{4, 4, 6, 1, 6, 7, 1, 7, 89, 21, 98}
	num2 := []int{1, 1, 3, 56, 1, 3, 5, 6, 1, 90, 6, 4, 7}
	fmt.Println("冒泡\n", num1)
	fmt.Println(paixu(s, maopaoSort))
	fmt.Println("选择\n", num2)
	fmt.Println(paixu(num2, xuanzeSort))
}

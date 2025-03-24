package main

import (
	"fmt"
	"strings"
)

func main() {
	lengthOfLongestSubstring("abba")
}
func lengthOfLongestSubstring(s string) int {
	str := strings.Split(s, "")
	var len1 int
	var maxLen int
	m := make([]string, 0)
	var left int
	var right int

	for i := range str {
		for k, v := range m {
			if v == str[i] {
				m = m[k+1:]
				break
			}
		}

		m = append(m, str[i])
		len1 = len(m)

		if i == 0 {
			maxLen = len1
		} else {
			maxLen = max(maxLen, len1)
		}

		fmt.Println(maxLen, len1, m, left, right)
	}
	return maxLen
}

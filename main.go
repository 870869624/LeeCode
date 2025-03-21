package main

import "fmt"

func main() {
	groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
}
func groupAnagrams(strs []string) [][]string {
	var bucket = make(map[[26]int][]string)

	for _, str := range strs {
		var count [26]int
		for _, c := range str {
			count[c-'a']++
		}
		bucket[count] = append(bucket[count], str)
	}

	var ret = make([][]string, 0, len(bucket))

	fmt.Println(bucket, "==========")
	for _, v := range bucket {
		ret = append(ret, v)
		fmt.Println(v, "-----", ret)
	}
	return
}

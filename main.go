package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(2))
}

func generateParenthesis(n int) []string {
	//每添加一个左括号，加1计数需要的右括号

	res := []string{}

	var dfs func(left, right int, str string)
	dfs = func(left, right int, str string) {
		if right > 0 {
			fmt.Println(left, right, "-------")
			dfs(left, right-1, str+")")
		}

		if left > 0 {
			fmt.Println(left, right, "+++++++")
			dfs(left-1, right+1, str+"(")
		}

		if left == 0 && right == 0 {
			res = append(res, str)
		}
	}
	dfs(n, 0, "")
	return res
}

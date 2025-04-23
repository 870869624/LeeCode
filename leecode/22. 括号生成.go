package main

func generateParenthesis(n int) []string {
	//每添加一个左括号，加1计数需要的右括号
	res := []string{}

	var def func(left int, right int, str string)

	def = func(left int, right int, str string) {
		if right > 0 {
			def(left, right-1, str+")")
		}
		if left > 0 {
			def(left-1, right+1, str+"(")
		}
		if left == 0 && right == 0 {
			res = append(res, str)
		}

	}
	def(n, 0, "")
	return res
}

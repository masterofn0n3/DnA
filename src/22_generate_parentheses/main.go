package main

import "fmt"

/*
()()()
(())()
choice: close paren or not open
*/
func generateParenthesis(n int) []string {
	result := []string{}

	var generate func(str string, open, close int)

	generate = func(str string, open, close int) {
		if len(str) == n*2 {
			result = append(result, str)
			return
		}
		if open > 0 {
			generate(str+"(", open-1, close)
		}
		if close > open {
			generate(str+")", open, close-1)
		}
	}

	generate("", n, n)

	return result
}

func main() {
	fmt.Println(generateParenthesis(3))
}

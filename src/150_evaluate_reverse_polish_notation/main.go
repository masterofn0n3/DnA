package main

import "strconv"

// ["2","1","+","3","*"]
func evalRPN(tokens []string) int {
	opeStack := []int{}
	for _, v := range tokens {
		n, err := strconv.Atoi(v)
		if err == nil {
			opeStack = append(opeStack, n)
		} else {
			v1 := opeStack[len(opeStack)-1]
			opeStack = opeStack[:len(opeStack)-1]
			v2 := opeStack[len(opeStack)-1]
			opeStack = opeStack[:len(opeStack)-1]
			if v == "+" {
				opeStack = append(opeStack, v1+v2)
			} else if v == "-" {
				opeStack = append(opeStack, v2-v1)
			} else if v == "*" {
				opeStack = append(opeStack, v1*v2)
			} else {
				opeStack = append(opeStack, v2/v1)
			}
		}
	}
	return opeStack[0]
}

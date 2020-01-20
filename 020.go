package main

import "fmt"

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	stack := make([]int32, len(s))
	index := 0
	for _, v := range s {
		if v == '(' || v == '[' ||v == '{' {
			stack[index] = v
			index ++
		} else {
			if index == 0 {
				return false
			}
			if v == ')' && stack[index - 1] == '(' {
				index --
			} else if v == ']' && stack[index - 1] == '[' {
				index --
			} else if v == '}' && stack[index - 1] == '{'{
				index --
			}  else {
				return false
			}
		}
	}
	return index == 0
}

func main()  {
	fmt.Println(isValid("()[]{}"))
}
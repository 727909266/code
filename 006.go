package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	res := make([]string, numRows)
	flag := -1
	row:=0
	for i:=0;i<len(s);i++ {
		res[row] += string(s[i])
		if row ==0 || row == numRows - 1 {
			flag = -flag
		}
		row += flag
	}

	return strings.Join(res, "")
}

func main()  {
	fmt.Println(convert("LEETCODEISHIRING", 1))
	fmt.Println(convert("LEETCODEISHIRING", 3))
	fmt.Println(convert("LEETCODEISHIRING", 4))
}
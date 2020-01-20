package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

func maxArea(height []int) int {
	res := 0
	l :=0
	r := len(height) -1
	for {
		if(l >= r) {
			break
		}
		res = max(res, (r-l) * min(height[l], height[r]))
		if height[l] > height[r] {
			r --
		} else {
			l ++
		}
	}

	return res
}

func main()  {
	fmt.Println(maxArea([]int{
		1,
		2,
	}))
}
package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	ans := 0
	for {
		if x == 0 {
			break
		}

		if ans>math.MaxInt32/10 || ans==math.MaxInt32/10&&x%10>7 { //判断上限
			return 0
		}
		if ans<math.MinInt32/10 || ans==math.MinInt32/10&&(x%10)<(-8) {
			return 0
		}


		ans *= 10
		ans += x%10
		x/=10
	}
	return ans
}

func main()  {
	fmt.Println(reverse(5678))
}
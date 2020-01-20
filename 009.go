package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 || x %10 ==0&& x !=0{
		return false
	}
	res := 0
	for {
		if res >= x {
			break
		}
		res = res * 10 + x % 10
		x /= 10
	}

	if x == res || x == res / 10 {
		return true
	}

	return false
}

func main()  {
	fmt.Println(isPalindrome(121))
}
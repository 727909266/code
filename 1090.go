package main

import "sort"

type ss struct {
	v int
	l int
}

func largestValsFromLabels(values []int, labels []int, num_wanted int, use_limit int) int {
	ans := 0
	ll := 0
	s := make([]ss, len(values))
	for i, _ := range  values {
		s[i].v = values[i]
		s[i].l = labels[i]
	}
	sort.Slice(s, func(i, j int) bool {
		if s[i].v > s[j].v{
			return true
		}
		return false
	})

	flag := make(map[int]int, 0)
	for i := 0; i < len(values); i ++ {
		if flag[s[i].l] < use_limit {
			flag[s[i].l] += 1
			ans += s[i].v
			ll ++
			if ll ==  num_wanted {
				return ans
			}
		}
	}

	return ans
}

func main()  {

}
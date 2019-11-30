package main

import (
	"container/list"
	"fmt"
	"sort"
)

func eventualSafeNodes(graph [][]int) []int {
	ans := make([]int, 0)
	graphc := make(map[int]int, 0)
	graphin := make(map[int][]int, 0)
	queue := list.New()
	for i, v := range graph {
		graphc[i] = len(v)
		if len(v) == 0 {
			queue.PushBack(i)
		} else {
			for _, j := range v {
				graphin[j] = append(graphin[j], i)
			}
		}
	}

	for {
		if queue.Len() == 0 {
			break
		}
		v := queue.Front().Value.(int)
		queue.Remove(queue.Front())
		fmt.Println(v)
		ans = append(ans, v)
		for _, vv := range graphin[v] {
			graphc[vv] --
			if graphc[vv] == 0 {
				queue.PushBack(vv)
			}
		}
	}


	sort.Ints(ans)
	return ans
}
func main()  {
	ss := [][]int {
		{},{0,2,3,4},{3},{4},{},
	}
	fmt.Println(eventualSafeNodes(ss))
}
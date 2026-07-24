package main

import "fmt"

func pairSumUnsorted(nums []int, target int) []int {
	hashmap := make(map[int]int)

	for i, x := range nums {
		complement := target - x
		if j, found := hashmap[complement]; found {
			return []int{j, i}

		}
		hashmap[x] = i
	}
	return []int{}
}

func exec_pairSumUnsorted() {
	fmt.Println(pairSumUnsorted([]int{1, 3, 5, 8}, 11))
}

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var res []int
	for i, v1 := range nums {
		for j, v2 := range nums {
			if i < j && v1 + v2 == target {
				res = append(res, v1, v2)
			}
		}
	}
	return res
}

func twoSum2(nums []int, target int) []int {
	var res []int
	var m map[int]int
	m = make(map[int]int)
	for i, v := range nums {
		fmt.Println(m[v])
		if m[v] != 0 && v + v == target {
			res = append(res, m[v]-1, i)
			break
		}
		m[v] = i+1
	}
	if (len(res) != 0) {
		return res
	}
	for i,v := range nums {
		if m[target-v] != 0 {
			res = append(res, i, m[target-v]-1)
            break
		}
	}
	return res
}

func main() {
	var r []int
	a := []int{2, 7, 11, 15}
	r = twoSum2(a, 9)
	fmt.Println(r)
}
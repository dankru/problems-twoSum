package main

import "fmt"

func main() {
	// O(n) Time and space
	fmt.Println(twoSum([]int{3, 3}, 6))
}

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	m[nums[0]] = 0
	for i := 1; i < len(nums); i++ {
		wantedNumber := target - nums[i]
		index, ok := m[wantedNumber]
		if ok {
			return []int{index, i}
		}
		m[nums[i]] = i
	}
	return []int{}
}

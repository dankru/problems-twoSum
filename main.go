package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {

	for i, v := range nums {
		for j, n := range nums[i+1:] {

			if v+n == target {
				return []int{i, j + i + 1}
			}
		}
	}
	return nil
}

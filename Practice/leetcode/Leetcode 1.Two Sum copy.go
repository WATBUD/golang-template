package main

import (
	"fmt"
)

func StartTwoSumcopy() {
	// //twoSum([]int{15, 7, 11, 2}, 9)
	_a := twoSum([]int{15, 7, 11, 2}, 9)
	fmt.Println(_a)
	//pointerSymbols()
}

func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{-1, -1}
}

func TwoSumHash(nums []int, target int) []int {

	_map := make(map[int]int)
	for index, value := range nums {

		j, exist := _map[target-value]
		fmt.Println(_map)
		if exist {
			return []int{index, j}
		}
		fmt.Println("_value:", value, "_index:", index)
		_map[value] = index
	}

	return []int{-1, -1}
}

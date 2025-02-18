package main

import "fmt"

func main() {

	fmt.Println(twoSum([]int{3, 3}, 6))

}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	var solution []int

	for idx, v := range nums {
		if value, exists := m[v]; exists && m[v] != idx {
			solution = append(solution, value, idx)
		}

		m[target-v] = idx

	}
	return solution
}

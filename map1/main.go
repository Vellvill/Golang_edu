package main

import "fmt"

func Interception(first, second []int) []int {
	result := make(map[int]int)
	array1 := []int{}
	for _, v := range first {
		result[v] = 1
	}
	for _, v := range second {
		num, ok := result[v]
		if ok {
			if num == 1 {
				array1 = append(array1, v)
				result[v]++
			}
		}
	}
	return array1
}

func twoSum(nums []int, target int) (int, int) {
	m := make(map[int]int)
	for i, v := range nums {
		if j, ok := m[v]; ok {
			return j, i
		}
		m[target-v] = i
	}
	return -1, -1
}

func main() {
	first := []int{0, 2, 4, 6, 8, 10}
	second := []int{3, 4, 5}

	fmt.Println(Interception(first, second))
	fmt.Println(twoSum(first, 18))
}

package main

import "fmt"

func sum(a []int) int {
	m := 0
	for _, v := range a {
		m += v
	}
	return m
}

func createSubsequence(a int) []int {
	massive := []int{}
	for i := 1; i <= a; i++ {
		massive = append(massive, i)
	}
	return massive
}

func main() {
	a, c := 100, 100
	cmassive := createSubsequence(c)
	cmassive[9] = 0
	fmt.Println(sum(createSubsequence(a)) - sum(cmassive))
}

package main

import "fmt"

//func findNumber(a []int) int {

}

func sum(a []int) int{
	m := 0
	for _, v := range a{
		m +=v
	}
	return m
}

func main() {
	a := []int{1,2,3}
fmt.Println(sum(a))
}







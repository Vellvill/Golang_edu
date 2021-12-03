package main

import (
	"fmt"
	"sort"
)

var res = []int{}

func main(){
	a := []int{6,6,3}
	b := []int{3,4,6}
	fmt.Println(interception(a,b))
}

func interception(x, y []int) ([]int){
	sort.Ints(x)
	sort.Ints(y)

	if len(x) == 0 || len(y) == 0{
		return res
	}
	var i, j int
	for {
		if x[i] == y[j]{
			res = append(res, x[i])
		}
		if j ++; j >= len(y){
			break
		}
		if i = sort.SearchInts(x, y[j]); i >= len(x){
			break
		}
	}
	return res


}


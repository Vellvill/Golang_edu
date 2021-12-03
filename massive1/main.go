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
	sort.Ints(x) //сорт для searchint
	sort.Ints(y)

	if len(x) == 0 || len(y) == 0{ //условие нулевого массива
		return res
	}
	var i, j int //перменные для цикла
	for {
		if x[i] == y[j]{ //пуш в пустой массив res
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


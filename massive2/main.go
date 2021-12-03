package main

import "fmt"

func main(){
	x:=[]int{1,2,3,7,4,4,4,4,5}
	y:=[]int{4,5,6}
	fmt.Println(Interception(x,y))
}

func Interception(x, y []int) []int{
	res := []int{}
	if len(x) == 0 || len(y) == 0{
		return res
	}

		for _, v:= range x {
			for _, z := range y {
				if v == z {
					m := v
					res = append(res, m)
				}
			}
		}

	var resNew[]int

	for _, v := range res {
		skip := false
		for _, u := range resNew {
			if v == u {
				skip = true
				break
			}
		}
		if !skip {
			resNew = append(resNew, v)
		}
	}

	return resNew
}


package main

import (
	"fmt"
)
var number int = 10

func main(){
	i := -10
	for a := true ; a ; a = i < 0 {
		fmt.Println(i)
		fmt.Println(i+number)
		i++
		if i == -1{
			break
		}
	}
}
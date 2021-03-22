package main

import (
	"fmt"
)

func main(){
	const count int = 10
	for i := 0;i < 10; i++{
		if i == 5 {
			continue
		}
		fmt.Println(count+i)

	}
}
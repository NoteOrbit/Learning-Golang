package main

import "fmt"

func main() {
	// variable
	// var + var_name + type_var
	var num int
	num = 10
	//var test string
	//test = "hello --> Note_Orbit"
	
	name := "hello -> "
	//Fl := 99.60
	const Fl float64 = 99.60 // constant variable
	var login string


	// fmt.Printf("Hello ~> go %b %v",num,name)
	s := fmt.Sprintf("%v / %v ",name,num)
	fmt.Printf("SF =>>",s)
	fmt.Scanf("%v \n",&login)
	// fmt.Printf(login)
	
	a := true //boolean --> test..

	if a{
		fmt.Printf("%v%v \n",name,login)
		for i :=0; i < 10; i++{
			result := Fl + float64(i) // operations same and type have to same
			fmt.Printf("%.2f\n",result)
		}
	} else {
		fmt.Printf("Dont know")
	}

	switch login {
	case "NoteOrbit":
		fmt.Printf("Hello --> %v",login)
	case "":
		fmt.Printf("Hello --> Dont know")
	}
} 
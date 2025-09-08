package main

import "fmt"

// Note: This is majorly used in idiomatic Go, for example to return both result and error values from a function
func vals()(int, int){
	return 3, 7 //The (int, int) in this function signature shows that the function returns 2 ints.
}
func MultipleReturnValues(){
	a, b := vals() //Here we use the 2 different return values from the call with multiple assignment.
	fmt.Println(a)
	fmt.Println(b) 

	_, c:= vals() //If you only want a subset of the returned values, use the blank identifier _.

	fmt.Println(c)
}
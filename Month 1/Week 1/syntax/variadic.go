package main

import "fmt"

func sum(nums ...int){ //Here’s a function that will take an arbitrary number of ints as arguments.
	fmt.Println(nums, " ")
	total:=0

	for _, num := range nums{
		total += num //Within the function, the type of nums is equivalent to []int. We can call len(nums), iterate over it with range, etc.
	}
	fmt.Println(total)
}
func Variadic(){
	sum(1, 2)
	sum(1, 2, 3) //Variadic functions can be called in the usual way with individual arguments.

	nums := []int{1, 2, 3, 4}
	sum(nums...) //If you already have multiple args in a slice, apply them to a variadic function using func(slice...) like this.
}
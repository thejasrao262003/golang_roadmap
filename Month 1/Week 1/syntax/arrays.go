package main

import "fmt"


//note [...]int{1, 2, 3, 4} -> This tells the compiler to count the number of elements and fixes the length at compile time. It can't be changed
// [] int{1, 2, 3, 4} -> similar to vectors in cpp. Can add or remove elements from array dynamically. It just creates a reference/pointer to array location
func Arrays(){
	var a[5] int
	fmt.Println(a)

	a[4] = 200

	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])
	fmt.Println("len: ", len(a))

	b := [5] int{1, 2, 3, 4, 5}
	// b = {1, 2, 3, 4, 5, 6}
	fmt.Println("dcl:", b)

	b = [...]int{1, 2, 3, 4, 5} //Use this syntax to declare and initialize an array in one line.
	fmt.Println("dcl:", b)

	b = [...]int{100, 3: 400, 500} //If you specify the index with :, the elements in between will be zeroed.
	fmt.Println("idx:", b)

	var twoD [2][3]int //Array types are one-dimensional, but you can compose types to build multi-dimensional data structures.
	for i:=0;i<len(twoD);i++{
		for j:=0;j<len(twoD[0]);j++{
			twoD[i][j] = i+j+1
		}
	}
	fmt.Println(twoD)
}
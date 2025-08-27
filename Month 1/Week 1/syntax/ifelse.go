package main

import "fmt"

// note: You don’t need parentheses around conditions in Go, but that the braces are required.
func IfElse() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else{
		fmt.Println("7 is odd")
	}

	if 8%2 == 0 || 7%2 == 0 { //Logical operators like && and || are often useful in conditions.
        fmt.Println("either 8 or 7 are even")
    }

	if num := 9; num<0 { //A statement can precede conditionals; any variables declared in this statement are available in the current and all subsequent branches.
		fmt.Println(num, "is negative")
	}else if num > 0{
		fmt.Println(num, "is positive")
	}else{
		fmt.Println(num, "is zero")
	}
}
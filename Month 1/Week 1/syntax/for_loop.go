package main

import "fmt"

func ForLoop(){
	i := 1
	fmt.Println("Using For Loop like while loop")
	for i <= 3{ //The most basic type, with a single condition. Similar to while loop
		fmt.Println(i)
		i += 1
	}
	fmt.Println("Using for loop")
	for j := 0; j<3; j++{ //A classic initial/condition/after for loop.
		fmt.Println(j)
	}

	fmt.Println("Using range for 'for loop'")
	for k := range 3{ //Another way of accomplishing the basic “do this N times” iteration is range over an integer.
		fmt.Println(k)
	}
	i = 0
	for{ //for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.
		fmt.Println("loop")
		i++
		if i>2 {
			break
		}
	}

	for n := range 6{ //You can also continue to the next iteration of the loop.
		if n%2 == 0{
			continue
		}
		fmt.Println(n)
	}
}
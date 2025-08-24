package main

import (
	"fmt"
	"example.com/greetings"
)
func main(){
	message := greetings.Hello("Thejas")
	fmt.Println(message)
}
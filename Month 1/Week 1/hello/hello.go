package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)
func main(){

	log.SetPrefix("Greetings: ")
	log.SetFlags(0)
	message, err := greetings.Hello("")
	if err != nil{
		log.Fatal(err)
	}else{
		fmt.Println(message)
	}
}
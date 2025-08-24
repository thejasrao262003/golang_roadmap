package greetings

import "fmt"

func Hello(s string) string {
	message := fmt.Sprintf("Hello, %v. Welcome!", s)
	return message
}
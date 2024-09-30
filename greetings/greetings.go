package greetings

import "fmt"

// hello returns a greeting for the name person
func Hello(name string) string {
	// return a greeting that embeds the name in a message
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

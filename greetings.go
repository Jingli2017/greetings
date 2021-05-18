package greetings

import "fmt"

// Hello version v2.0.0
// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Holla, %v. Welcome to Golang!", name)
	return message
}

package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello version v2.0.0
// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == ""{
		return "", errors.New("empty name")
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := map[string]string{}
	for _, name := range names{
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	options := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return options[rand.Intn(len(options))]
}

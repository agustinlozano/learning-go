package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // Return a greeting that embeds the name in a message.

		if name == "" {
			return "", errors.New("empty name")
		}

    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

// Hellos returns a map that associates each of the named
// people with a greeting message
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	// make(map[key-type]value-type
	messages := make(map[string]string)

	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string {
		"Hi, %v. Welcome",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}

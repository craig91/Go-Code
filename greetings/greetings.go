package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("empy name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	// message := fmt.Sprintf(randomFormat()) <-- bad code for testing if error handling works via the go test.
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {
	// A slice of message formats
	formats := []string{
		"Hi, %v. Welcome! ",
		"Great to see you, %v ",
		"Hail, %v! Well met! ",
	}

	return formats[rand.Intn(len(formats))]
}

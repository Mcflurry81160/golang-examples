package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("no name specified")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	//This is a key value pair essentially
	messages := make(map[string]string)

	//'Range' iterates through the members in the for loop
	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		//adding an entry to a map
		messages[name] = message
	}

	return messages, nil
}

func init() {
	//This essentially runs before the main function in the package is run.
	//See https://stackoverflow.com/a/24790378/1369341 for a good explanation of this
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	//Must have a comma at the end of each entry
	formats := []string{
		"Hi, %v!",
		"Hi, %v! Good day!",
	}

	return formats[rand.Intn(len(formats))]
}

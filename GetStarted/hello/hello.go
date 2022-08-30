package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {

	names := []string{"someName"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	//This will print the map as it is
	fmt.Println(messages)

	//One way of accessing a specific entry's value in a map
	fmt.Println(messages["someName"])

	//Another way of accessing a specific entry's value in a map
	fmt.Println(messages[names[0]])
}

package greetings

import (
	"regexp"
)

func TestHelloName() {
	name := "SomeName"
	want := regexp.MustCompile(name)
}

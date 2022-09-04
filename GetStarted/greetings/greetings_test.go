package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "SomeName"
	want := regexp.MustCompile(name)

	message, err := Hello("SomeName")

	if want.MatchString(message) && err == nil {
		return
	}

	t.Fatalf(`Hello("SomeName") = %q, %v, want match for %#q, nil`, message, err, want)
}

func TestHelloEmpty(t *testing.T) {
	message, err := Hello("")

	//This is testing if the error handling works in the Hello function.
	if message == "" && err != nil {
		return
	}

	t.Fatalf(`Hello("") = %q, %v, want "", error`, message, err)
}

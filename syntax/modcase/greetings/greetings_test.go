package greetings

import (
	"regexp"
	"testing"
)

func TestHello(t *testing.T) {
	name := "Tom"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Tom")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Tom") = %q, %v, want %q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

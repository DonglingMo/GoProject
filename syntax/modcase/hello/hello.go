package main

import (
	"GoProject/syntax/modcase/greetings"
	"fmt"
	"log"
)

import "rsc.io/quote"

func main() {
	// go mod tidy case
	fmt.Println(quote.Go())
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	names := []string{"Alice", "Bob", "Cecil", "Dave"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}

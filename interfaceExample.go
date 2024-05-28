package main

import "fmt"

type IBot interface {
	getGreeting() string
}

type EnglishBot struct{}

type SpanishBot struct{}

func (eb EnglishBot) getGreeting() string {
	// Custom Logic
	return "Hi There"
}

func (sb SpanishBot) getGreeting() string {
	return "Hola"
}

func printGreeting(b IBot) {
	fmt.Println(b.getGreeting())
}

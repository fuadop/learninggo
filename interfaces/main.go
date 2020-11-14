package main

import "fmt"

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

type bot interface {
	getGreeting() string
}

type englishBot struct{}

type spanishBot struct{}

func (englishBot) getGreeting() string {
	return "Hello There"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

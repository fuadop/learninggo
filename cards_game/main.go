package main

import "fmt"

func main() {
	cards := newDeck()
	cards.saveToFile("./mycards.txt")

	fmt.Println(cards.toString())
}

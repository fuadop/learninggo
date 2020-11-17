package main

import "fmt"

func main() {
	var i integer = 5
	i.ToString()
	fmt.Println(i)
}

type integer int

func (i *integer) ToString() {
	(*i) += 5
}

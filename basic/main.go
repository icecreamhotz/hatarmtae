package main

import "fmt"

func main() {
	var cards []int
	for i := 0; i < 50; i++ {
		cards = append(cards, i)
	}
	fmt.Println(cards)
}

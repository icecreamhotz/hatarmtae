package main

import "fmt"

func main() {
	for i := 1; i <= 50; i += 1 {
		if i%2 == 0 {
			fmt.Printf("%d : %s\n", i, "even")
		} else {
			fmt.Printf("%d : %s\n", i, "odd")
		}
	}
}

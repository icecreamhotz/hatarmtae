package main

import "fmt"

func main() {
	lnwza := make(map[string]int)
	lnwza["joe"] = 112
	lnwza["karn"] = 29
	for k, v := range lnwza {
		fmt.Printf("%s : %d\n", k, v)
	}
}

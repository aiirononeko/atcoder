package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scanf("%s", &s)
	var count int
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "1" {
			count++
		}
	}
	fmt.Print(count)
}

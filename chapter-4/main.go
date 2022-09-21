package main

import "fmt"

// If statements
// for loop

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("number", i, "even")
		} else {
			fmt.Println("number", i, "odd")
		}
	}
}

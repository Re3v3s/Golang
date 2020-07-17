package main

import (
	"fmt"
)

//! main() first func to call . it's like __contruct function
func main() {
	var lim int
	fmt.Printf("Limit: ")
	fmt.Scanf("%d", &lim) //&lim = move value to lim
	// short if statement
	if v := 2; v > lim {
		fmt.Println("Limit < 2")
	} else {
		fmt.Println("Limit > 2")
	}
}

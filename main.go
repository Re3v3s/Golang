package main

import (
	"fmt"
)

//! main() first func to call . it's like __contruct function
func main() {
	var n1 int  = 5 
	n2:= 6
	n3,n4,n5 := true ,false ,"OK"

	fmt.Println(add(n1, n2))

	fmt.Println("test character n boolean : ", n3,n4,n5)
}

func add(x, y int) int {
	return (x + y)
}

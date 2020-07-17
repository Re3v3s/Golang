package main

import (
	"fmt"
)

func main() {
	i := 42
	p := &i //&i make addres of i to p  it's mean p = pointer

	fmt.Println("I = ", *p) // value of P  " * " mean value of P
	fmt.Println("Adress of I = ", p)

	*p = *p / 2

	fmt.Println("I = ", *p)
}

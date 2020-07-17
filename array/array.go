package main

import (
	"fmt"
)

func  main()  {
	var a[2] string
	a[0] = "Zero"
	a[1] = "One"
	fmt.Println(a)
	fmt.Println("1.",a[0] ,"2.",a[1])

	// short course
	course := [4]string{"PHP","JavaScript","ajax","Laravel"}
	fmt.Println(course)

	course2 := []string{"PHP","JavaScript","ajax","Laravel"}
	fmt.Println(course2)

	primes := [6]int {2,3,5,7,11,13}
	i := 4
	var s []int = primes[1:i]   // cut index 1:to i it's mean 1: 4 {2 3 5 7 11 13 } it's show  3 5 7 not count index [4]
	fmt.Println(s)

}
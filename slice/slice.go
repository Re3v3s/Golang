package main

import (
	"fmt"
)
// ! slice and decrease and addition 
func  main()  {
	primes := [6] int {1,2,3,4,5,6}
	var s []int = primes[3:5]
	fmt.Println(s)

	slicePrimes := []int {2,3,4,5,6,7,8,9}
	printSlices(slicePrimes)

	printSlices(slicePrimes[2:])  //slice only Leading
	printSlices(slicePrimes[:5])  //slice only Tailing
	printSlices(slicePrimes[3:6])


	slicePrimes = append(slicePrimes , 11)
	printSlices(slicePrimes)

	slicePrimes = append(slicePrimes , 13)
	printSlices(slicePrimes)

	slicePrimes = append(slicePrimes , 15,17)
	printSlices(slicePrimes)


}

func printSlices(s []int)  {
	fmt.Printf("len=%d cap=%d %v\n ",len(s) , cap(s) , s)
	// cap =  capacity = ปริมาณ ความจุ
	// len = length = ความยาว
}
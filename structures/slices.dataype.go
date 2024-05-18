package structures

import "fmt"

func WhatASliceSays() {
	fmt.Print("Slices: ")
	someSlice := []int{1} //this guy right here is also an array, but in Go, we'll called it 'Slices'

	/*
	A slice it's used when you didn't know how many positions inside an array you'll need.
	As you can't change the size of an array after created, this approach can supply you to interact with dinamic values

	Under the hood, a slice is an array, that will recieve the size of the values you passed when created it

	When you need more spaces in your slice, the go you'll give more spaces considering the size you passed when the slice was created.

	*/

	fmt.Print(someSlice)

	someSlice = append(someSlice, 2, 3) // the append method adds values to your slice and save this to a variable. In this case, I used the same.


	fmt.Println(someSlice)
}
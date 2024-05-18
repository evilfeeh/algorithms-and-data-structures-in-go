package structures

import "fmt"

// structs are like objects in Javascript, you can use it for join diferent data types
// it's possible to use for build a user, for example
type Person struct {
	name string
	age int
	handsome bool
}


func WhatTheStruct() {
	fmt.Print("Structs: ")
	person := Person{name: "Felipe", age: 18, handsome:true}
	/*
	this is one possibility to implement a struct, but you can also ommit the keys, just passing the values, like that:
	persont := Person{"John", "30", false}
	*/
	fmt.Println(person)
}
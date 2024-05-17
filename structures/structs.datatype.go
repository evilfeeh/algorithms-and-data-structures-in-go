package structures

import "fmt"

type Person struct {
	name string
	age int
	handsome bool
}


func WhatTheStruct() {
	person := Person{name: "Felipe", age: 18, handsome:true}
	fmt.Println(person)
}
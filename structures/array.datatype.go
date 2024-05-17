package structures

import "fmt"

func Foods()	{
	food := [3]string{"pizza", "tacos", "sushi"} //this is an array.
	quantity := [3]int{1, 5, 10} // this too
	matrix := [2][2]int{{1,2},{3,4}} // I need to talk about this one?
	/*
	An array its a collection of items of the same type.
	You can create an array of strings, integers, floats and even arrays.

	It's very useful for ordered lists, some types of data and you can iteract by each item of an array.
	*/
	fmt.Println(food, quantity, matrix)
}
package structures

import "fmt"

func GiveMeSomeNumbers() {
	integerOne := 3 //this is a integer
	floatOne := 1.5 // this is a float64
	/*
	Numbers are self-explained, you can use for a lot of things.
	Actually, everything is number under the hood.

	In go, you can have different types of number, like:
	- integers
	-floats (this part can be float32 or float64)

	You can especify the type of float you want to use, this way:
	*/
	var lilFloat float32 = 1.4 

	fmt.Println(integerOne, floatOne, lilFloat)
}
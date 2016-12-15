package main

import "fmt"

func main() {

	msg := "Pointers are a type of variable that are going to contain the address of another variable in memory of the same type of the pointer"

	msg2 := "For example when we pass a variable to a func we need to understand how that variable is going to behave"
	var pointersExplained *string = &msg

	fmt.Println(msg, pointersExplained, msg2)
	msg3 := "When we put a star in front of the pointer variable, is going to reference the memory adreess and get the value. Ex:"

	fmt.Println(msg3, *pointersExplained)
	msg4 := "Now we change the value by making an assigment"

	*pointersExplained = "New Value"
	fmt.Println(msg4, *pointersExplained, "--This is the initial var:", msg)

}

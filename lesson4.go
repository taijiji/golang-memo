// pointer: variable to designate machine address
// Cannot caluculate.
package main

import "fmt"

func main() {
	a := 5
	var pointer_a *int
	pointer_a = &a 
	// &a 			: address of a
	// *pointer_a 	: data value of pointer_a 
	fmt.Println(pointer_a)
	fmt.Println(*pointer_a)
}
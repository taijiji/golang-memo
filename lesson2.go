/*
Initinal value

var s string // "". not nil.
var a int // 0
var b bool // false
*/
package main

import "fmt"

func main() {
	a := 10
	b := 12.3
	c := "text"
	var d bool
	fmt.Printf("a: %d, b:%f, c:%s, d:%t\n", a, b, c, d)

	fmt.Printf("a: %v, b:%v, c:%v, d:%v\n", a, b, c, d)
}
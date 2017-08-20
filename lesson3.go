// Constant

package main

import "fmt"

func main() {
	//const name = "constant"
	//name = "Taiji"
	//fmt.Println(name)
	//Error: "./lesson3.go:9: cannot assign to name"

	const (
		sun = iota // 0
		mon // 1
		tue // 2
	)
	fmt.Println(sun, mon, tue)
}
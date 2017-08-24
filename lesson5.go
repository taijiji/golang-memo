// fundcion
package main

import "fmt"

func hi(name string) {
	//fmt.Println("hi! " + name)
	message := "hi! " + name
	return message
}
func main() {
	//hi()
	//hi("Taiji")
	fmt.Println(hi("Taiji"))
}
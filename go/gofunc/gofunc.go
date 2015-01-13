// annonymous function and goruntime

package main

import (
	"fmt"
)

func main() {

	func() {
		fmt.Printf("hello\n")
	}()

	i := 3
	c := make(chan int)
	var result int
	go func() {
		result = i * 10
		c<- 1
	}()
	<-c
	fmt.Printf("%d\n",result)
}

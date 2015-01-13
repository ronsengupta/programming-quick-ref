/*
$ go run defer.go
0
4
3
2
1
*/

package main

import "fmt"

func f() {
	fmt.Printf("0\n")
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Printf("4\n")
}

func main() {
	f()
}

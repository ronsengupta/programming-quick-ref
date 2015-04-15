package main

import "fmt"

var (
	abc = tmpFunc()
)

func tmpFunc() int {
	fmt.Println("From var \n")
	return 0
}

func init() {
	fmt.Println("From init() \n")
}

func main() {
	fmt.Println("From main()\n")
}

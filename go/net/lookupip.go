package main

import (
	"fmt"
	"net"
)

func main() {
	ip, err := net.LookupIP("localhost")
	fmt.Printf("%v %v\n", ip, err)
}

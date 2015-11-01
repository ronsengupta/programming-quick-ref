package main

import (
	"fmt"
	"net"
)

func main() {
	intfs, err := net.Interfaces()
	if err != nil {
		fmt.Print(err)
		return
	}
	for _, i := range intfs {
		fmt.Printf("Name: %v\n", i.Name)
		fmt.Printf("Flags: %v\n", i.Flags)
	}

}

package main

import "fmt"

type tanto struct {
}

func (t *tanto) standup() {
	fmt.Printf("Tanto standup\n")
}

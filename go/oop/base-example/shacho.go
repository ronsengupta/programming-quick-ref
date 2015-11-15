package main

import (
	"os"
)

func main() {
	if len(os.Args) == 1 {
		os.Exit(0)
	}

	if os.Args[1] == "tanto" {
		var s Shain = &tanto{}
		s.standup()
	} else if os.Args[1] == "shunin" {
		var s Shain = &shunin{}
		s.standup()
	} else {
		os.Exit(0)
	}
}

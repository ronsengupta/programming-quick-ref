package main

import "fmt"

func main() {

	var timeZone = map[string]int{
		"UTC": 0,
		"EST": -5,
		"CST": -6,
		"MST": -7,
		"PST": -8,
	}

	fmt.Printf("%d\n", timeZone["EST"])

	timeZone["JST"] = +9

	fmt.Printf("%d\n", timeZone["JST"])

}

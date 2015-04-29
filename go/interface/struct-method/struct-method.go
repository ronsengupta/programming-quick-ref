package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

func(p *Person) Name() string {
	return p.FirstName + " " + p.LastName
}

func main() {
	person := &Person{"Taro", "Yamada"}
	fmt.Printf(person.Name())
}

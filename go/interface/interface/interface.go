package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

type Named interface {
	Name() string
}

func (p *Person) Name() string {
	return p.FirstName + " " + p.LastName
}

func printName(named Named) {
	fmt.Println(named.Name())
}

func main() {
	person := &Person{"Taro", "Yamada"}
	printName(person)
}

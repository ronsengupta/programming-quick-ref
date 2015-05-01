package main

import fmt "fmt"

type Test struct {
	Name string
}

func (t Test) hello() {
	fmt.Printf("Hello, %s\n", t.Name)
}

func (t *Test) goodby() {
	fmt.Printf("Goodby, %s\n", t.Name)
}

func main() {
	t := Test{Name: "nak3"}

	t.hello()
	t.goodby()

	f_hello := t.hello
	f_goodby := t.goodby

	t.Name = "kenjiro"

	f_hello()
	f_goodby()
}

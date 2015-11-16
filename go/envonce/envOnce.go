package main

import (
	"fmt"
	"os"
	"sync"
)

var (
	foobar = &envOnce{
		names: []string{"foo", "bar"},
	}
)

type envOnce struct {
	names []string
	once  sync.Once
	val   string
}

func (e *envOnce) Get() string {
	e.once.Do(e.init)
	return e.val
}

func (e *envOnce) init() {
	for _, n := range e.names {
		fmt.Printf("-- %s -- \n", n)
		e.val = os.Getenv(n)
		if e.val != "" {
			return
		}
	}
}

func main() {
	fmt.Printf("value =  %s \n", foobar.Get())
	fmt.Printf("value =  %s \n", foobar.Get())
	fmt.Printf("value =  %s \n", foobar.Get())
	fmt.Printf("value =  %s \n", foobar.Get())
}

package main

import (
    "fmt"
    "log"
    "net"
)

func main() {
    fmt.Println("==== Test with [::1] === ")
    l, err := net.Listen("tcp6", "[::1]:12345")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(l.Addr())
    l.Close()

    fmt.Println("==== Test with localhost === ")
    l_, err := net.Listen("tcp6", "localhost:12345")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(l_.Addr())
    l.Close()
}


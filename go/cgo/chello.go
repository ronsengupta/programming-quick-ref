package main

// #include "hello.h"
import "C"

import (
        "fmt"
)

func main() {
	C.hello()
}

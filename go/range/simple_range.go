package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}

	for i, value := range nums {
		fmt.Printf("%d: value = %d\n", i, value)
	}

}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

func jsonmap() {
	inf, err := os.Open("./test.json")
	if err != nil {
		return
	}

	raw_data, err2 := ioutil.ReadAll(inf)
	if err2 != nil {
		fmt.Printf("err1\n")
	}

	decoded := map[string]interface{}{}
	err = json.Unmarshal(raw_data, &decoded)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("T -> ", reflect.TypeOf(decoded["T"]))
	fmt.Println("O -> ", reflect.TypeOf(decoded["O"]))
	fmt.Println("N -> ", reflect.TypeOf(decoded["N"]))

	keys := [3]string{"T", "O", "N"}

	for i := 0; i < len(keys); i++ {
		fmt.Println("\n--- test ---\n")
		switch decoded[keys[i]].(type) {
		case string:
			fmt.Println("decode[\"O\"]")
			fmt.Println(decoded["O"])
		case map[string]interface{}:
			fmt.Println("decode[\"T\"]")
			fmt.Println(decoded["T"])
		case []interface{}:
			fmt.Println("decode[\"N\"]")
			fmt.Println(decoded["N"])
		default:
			fmt.Println("default")
		}
	}
}

func main() {
	jsonmap()
}

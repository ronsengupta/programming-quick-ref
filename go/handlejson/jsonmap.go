package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

func decoder() (string, map[string]interface{}, error) {
	inf, err := os.Open("./v1.json")
	if err != nil {
		return "", nil, err
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

	fmt.Println("type test -> ", reflect.TypeOf(decoded["T"]))
	keys := "models"
	return keys, decoded, nil
}

func valueOutputs(keys string, decoded map[string]interface{}) {

	switch decoded[keys].(type) {
	case string:
		fmt.Println(decoded[keys])
	case map[string]interface{}:
		fmt.Println("\ngo loop ->\n")
		oldkeys := keys
		for f := range decoded[keys].(map[string]interface{}) {
			keys = f
		}
		fmt.Println("--keys -> ", keys)
		valueOutputs(keys, decoded[oldkeys].(map[string]interface{}))
	case []interface{}:
		fmt.Println("\ngo loop ->\n")
		fmt.Println(decoded[keys])
	default:
		fmt.Println("default")
	}
}

func main() {
	keys, decoded, _ := decoder()
	valueOutputs(keys, decoded)
}

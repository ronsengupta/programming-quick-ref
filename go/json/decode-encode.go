package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Msg struct {
	Msg string `json:"msg"`
	Id  int64  `json:"id"`
}

func WithTags() {
	raw_data := []byte(`{"msg": "Hello Go!", "id": 12345}`)
	decoded := Msg{}
	err := json.Unmarshal(raw_data, &decoded)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(decoded.Msg, decoded.Id)
	encoded, _ := json.Marshal(&decoded)

	//revert it to raw_data
	fmt.Println(string(encoded))
}

func main() {
	WithTags()
}

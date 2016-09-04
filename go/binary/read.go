package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
)

type Row struct {
	Key uint16
	Val uint16
}

func (r Row) String() string {
	return fmt.Sprintf("(%s: %v)", string(r.Key), r.Val)
}

func main() {
	buf := bytes.NewBuffer([]byte{0x00, 98, 0x0, 0x2})
	row := Row{}
	binary.Read(buf, binary.BigEndian, &row.Key)
	binary.Read(buf, binary.BigEndian, &row.Val)

	log.Println(row)
}

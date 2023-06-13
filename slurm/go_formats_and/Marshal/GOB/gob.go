package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

func GobExample() {
	DecodeGob(EncodeGob())
}

func EncodeGob() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	m := make(map[string]string)
	m["foo"] = "bar"

	if err := enc.Encode(m); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.Bytes()) //[14 255 129 4 1 2 255 130 0 1 12 1 12 0 0 12 255 130 0 1 3 102 111 111 3 98 97 114]

	return buf.Bytes()
}

func DecodeGob(input []byte) {
	buf := bytes.NewBuffer(input)
	dec := gob.NewDecoder(buf)

	m := make(map[string]string)
	if err := dec.Decode(&m); err != nil {
		log.Fatal(err)
	}

	fmt.Println(m["foo"])
}

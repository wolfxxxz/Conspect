package main

import (
	"Protobuf/Protobuf/gen"
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"
)

func ProtobufCase() {

	elliot := gen.Person{
		Name: "Elliot",
		Age:  24,
	}

	elliot.ProtoMessage()
	// Не работает
	data, err := proto.Marshal(elliot)
	if err != nil {
		log.Fatal("Marshalling error")
	}
	fmt.Println(data)
	newElliot := &gen.Person{}
	err = proto.Unmarshal(data, newElliot)
	if err != nil {
		log.Fatal("unmarshalling error", err)
	}

	fmt.Println(newElliot.GetAge())
	fmt.Println(newElliot.GetName())
}

package main

import (
	"fmt"

	"github.com/simplesteph/protobuf-example-go/src/simple"
)

func main() {
	doSimple()
}

func doSimple() {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 4, 7, 8},
	}
	fmt.Println(sm)

	sm.Name = "I renamed you"
	fmt.Println(sm)

	fmt.Println("The ID is:", sm.GetId())

}

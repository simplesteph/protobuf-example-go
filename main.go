package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/simplesteph/protobuf-example-go/src/complex"
	"github.com/simplesteph/protobuf-example-go/src/enum_example"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/simplesteph/protobuf-example-go/src/simple"
)

func main() {
	sm := doSimple()

	readAndWriteDemo(sm)
	jsonDemo(sm)

	doEnum()

	doComplex()
}

func doComplex() {
	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "First message",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Id:   2,
				Name: "Second message",
			},
			&complexpb.DummyMessage{
				Id:   3,
				Name: "Third message",
			},
		},
	}

	fmt.Println(cm)
}

func doEnum() {
	em := enumpb.EnumMessage{
		Id:           42,
		DayOfTheWeek: enumpb.DayOfTheWeek_THURSDAY,
	}

	em.DayOfTheWeek = enumpb.DayOfTheWeek_MONDAY
	fmt.Println(em)
}

func jsonDemo(sm proto.Message) {
	smAsString := toJSON(sm)
	fmt.Println(smAsString)

	sm2 := &simplepb.SimpleMessage{}
	fromJSON(smAsString, sm2)
	fmt.Println("Successfully created proto struct:", sm2)
}

func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
		return ""
	}
	return out
}

func fromJSON(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("Couldn't unmarshal the JSON into the pb struct", err)
	}
}

func readAndWriteDemo(sm proto.Message) {
	writeToFile("simple.bin", sm)
	sm2 := &simplepb.SimpleMessage{}
	readFromFile("simple.bin", sm2)
	fmt.Println("Read the content:", sm2)
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialise to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}

	fmt.Println("Data has been written!")
	return nil
}

func readFromFile(fname string, pb proto.Message) error {

	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Something went wrong when reading the file", err)
		return err
	}

	err2 := proto.Unmarshal(in, pb)
	if err2 != nil {
		log.Fatalln("Couldn't put the bytes into the protocol buffers struct", err2)
		return err2
	}

	return nil
}

func doSimple() *simplepb.SimpleMessage {
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

	return &sm
}

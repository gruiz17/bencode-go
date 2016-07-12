package bencode

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestDecodeString(t *testing.T) {
	str1 := "blahblah4:blahfasdf"
	blah, eIndex, err := readStr(8, str1)
	if err != nil {
		t.Error("something went wrong decoding string")
		log.Fatal(err)
	}
	if blah != "blah" {
		t.Error("string unequal")
	}
	if eIndex != 14 {
		t.Error("index is not right")
	}
}

func TestDecodeInt(t *testing.T) {
	int1 := "blahblahi424easdf"
	blah, eIndex, err := readInt(8, int1)
	if err != nil {
		t.Error("something went wrong decoding string")
		log.Fatal(err)
	} else if blah != 424 {
		t.Error("int unequal")
	} else if eIndex != 13 {
		t.Error("index is not right")
	}
}

func TestDecodeList(t *testing.T) {
	list1 := "l4:spam4:eggse"
	trueList := []interface{}{"spam", "eggs"}

	decodedList, eIndex, err := readColl(0, list1, "l")

	if err != nil {
		t.Error("something went wrong decoding list")
		log.Fatal(err)
	} else if !(reflect.DeepEqual(decodedList, trueList)) {
		fmt.Println("you got ", decodedList)
		fmt.Println("should be ", trueList)
		t.Error("lists not equal :O")
	} else if eIndex != 14 {
		t.Error("end index unequal nooooo")
	}
}

func TestDecodeNestedList(t *testing.T) {
	nested := "l4:spam4:eggsl4:spam4:eggsee"
	trueList := []interface{}{"spam", "eggs"}
	trueList = append(trueList, []interface{}{"spam", "eggs"})

	decodedList, eIndex, err := readColl(0, nested, "l")

	if err != nil {
		t.Error("something went wrong decoding nested list")
		log.Fatal(err)
	} else if !(reflect.DeepEqual(decodedList, trueList)) {
		fmt.Println("you got ", decodedList)
		fmt.Println("should be ", trueList)
		t.Error("nested lists not equal nooooo")
	} else if eIndex != 28 {
		t.Error("end index on nested list unequal nooooo")
	}
}

func TestDecodeDict(t *testing.T) {
	dict1 := "d3:cow3:moo4:spam4:eggseffff"
	trueDict := make(map[string]interface{})
	trueDict["cow"] = "moo"
	trueDict["spam"] = "eggs"
	decodedDict, eIndex, err := readColl(0, dict1, "d")
	if err != nil {
		t.Error("something went wrong decoding dict")
		log.Fatal(err)
	} else if !(reflect.DeepEqual(decodedDict, trueDict)) {
		fmt.Println("you got ", decodedDict)
		fmt.Println("should be ", trueDict)
		t.Error("dicts not equal")
	} else if eIndex != 24 {
		t.Error("end index unequal nooooo")
	}
}

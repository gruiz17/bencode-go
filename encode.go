package bencode

import (
	"reflect"
	"strconv"
	"unicode/utf8"
)

func encodeString(s string) string {
	length := strconv.Itoa(utf8.RuneCountInString(s))
	return length + ":" + s
}

func encodeInt(i int) string {
	return "i" + strconv.Itoa(i) + "e"
}

func encodeDict(d map[string]interface{}) string {
	returnedStr := "d"
	for k, v := range d {
		returnedStr += encodeString(k)
		switch reflect.TypeOf(v).Name() {
		case "string":
			returnedStr += encodeString(v.(string))
		case "int":
			returnedStr += encodeInt(v.(int))
		case "map[string]interface {}":
			returnedStr += encodeDict(v.(map[string]interface{}))
		case "[]interface {}":
			returnedStr += encodeList(v.([]interface{}))
		}
	}
	return returnedStr + "e"
}

func encodeList(l []interface{}) string {
	returnedStr := "l"
	for _, v := range l {
		switch reflect.TypeOf(v).Name() {
		case "string":
			returnedStr += encodeString(v.(string))
		case "int":
			returnedStr += encodeInt(v.(int))
		case "map[string]interface {}":
			returnedStr += encodeDict(v.(map[string]interface{}))
		case "[]interface {}":
			returnedStr += encodeList(v.([]interface{}))
		}
	}
	return returnedStr + "e"
}

package bencode

import (
	"testing"
)

func TestEncodeAllTheThings(t *testing.T) {
	d := make(map[string]interface{})
	e := []interface{}{}
	e = append(e, 1)
	e = append(e, 2)
	d["haha"] = "lol"
	d["lol"] = 20

	dEncode := "d3:loli20e4:haha3:lole"
	dEncodeAlt := "d4:haha3:lol3:loli20ee"
	dPossible := [2]string{dEncode, dEncodeAlt}

	eEncode := "li1ei2ee"

	dTest := encodeDict(d)
	eTest := encodeList(e)

	if eTest != eEncode {
		t.Error("problem encoding list")
	}

	somethingEqual := false
	for _, v := range dPossible {
		if dTest == v {
			somethingEqual = true
		}
	}
	if !(somethingEqual) {
		t.Error("problem encoding dict")
	}
}

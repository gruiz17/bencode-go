package bencode

import (
	"bytes"
	"errors"
	"regexp"
	"strconv"
)

func readColl(i int, encoded string, flag string) (interface{}, int, error) {
	var returnedColl interface{}
	returnedMap := make(map[string]interface{})
	returnedList := []interface{}{}
	switch flag {
	case "d":
		returnedList = nil
	case "l":
		returnedMap = nil
	}
	eIndex := i + 1

	numRegex, _ := regexp.Compile("[1-9]")

	// do this stuff if this coll is a dictionary
	currentKey := ""
	currentKeyFound := false

	// ok back to generic

	for string(encoded[eIndex]) != "e" {
		var newEIndex int
		var strFound string
		var err error

		if numRegex.MatchString(string(encoded[eIndex])) {
			strFound, newEIndex, err = readStr(eIndex, encoded)
			if err != nil {
				return nil, 0, err
			}
			switch flag {
			case "d":
				if currentKeyFound == false {
					currentKey = strFound
					currentKeyFound = true
				} else {
					returnedMap[currentKey] = strFound
					currentKey = ""
					currentKeyFound = false
				}
			case "l":
				returnedList = append(returnedList, strFound)
			}

		} else {
			var currVal interface{}
			switch string(encoded[eIndex]) {
			case "i":
				currVal, newEIndex, err = readInt(eIndex, encoded)
			case "l":
				currVal, newEIndex, err = readColl(eIndex, encoded, "l")
			case "d":
				currVal, newEIndex, err = readColl(eIndex, encoded, "d")
			default:
				return nil, 0, errors.New("PANIC PANIC PANIC")
			}
			switch flag {
			case "d":
				returnedMap[currentKey] = currVal
				currentKeyFound = false
			case "l":
				returnedList = append(returnedList, currVal)
			}
		}
		eIndex = newEIndex
	}
	switch flag {
	case "d":
		returnedColl = returnedMap
	case "l":
		returnedColl = returnedList
	}
	return returnedColl, eIndex + 1, nil
}

func readStr(i int, encoded string) (string, int, error) {
	lIndex := i
	var strLength bytes.Buffer
	for string(encoded[lIndex]) != ":" {
		strLength.WriteString(string(encoded[lIndex]))
		lIndex++
	}
	parsedLen, err := strconv.Atoi(strLength.String())
	if err != nil {
		return "", 0, err
	}
	sIndex := lIndex + 1
	eIndex := sIndex + parsedLen
	return encoded[sIndex:eIndex], eIndex, nil
}

func readInt(i int, encoded string) (int, int, error) {
	eIndex := i
	for string(encoded[eIndex]) != "e" {
		eIndex++
	}
	sIndex := i + 1
	parsedInt, err := strconv.Atoi(encoded[sIndex:eIndex])
	if err != nil {
		return 0, 0, err
	}
	return parsedInt, eIndex + 1, nil
}

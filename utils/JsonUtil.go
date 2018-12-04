package utils

import (
	"encoding/json"
	"log"
)

func MustToJsonString(val interface{}) string {
	var (
		result string
	)

	result, err := ToJsonString(val)
	if nil != err {
		log.Panic(err)
	}

	return result
}

func ToJsonString(val interface{}) (string, error) {
	var (
		result string
		err    error
	)

	data, err := json.Marshal(val)
	if nil != err {
		return result, err
	}

	result = string(data)

	return result, err
}

func ToObject(jsonString string, val interface{}) error {
	var (
		err error
	)

	err = json.Unmarshal([]byte(jsonString), &val)

	return err
}

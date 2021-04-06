package util

import "encoding/json"

//JSONMarshal Marshal json
func JSONMarshal(i interface{}) string {
	res, _ := JSONMarshalE(i)
	return res
}

//JSONMarshalE Marshal json
func JSONMarshalE(i interface{}) (string, error) {
	res, e := json.Marshal(i)
	if e != nil {
		return "", e
	}
	return string(res), nil
}

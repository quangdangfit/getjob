package utils

import (
	jsoniter "github.com/json-iterator/go"
)

// JSONToString JSON
func JSONToString(v interface{}) string {
	s, err := jsoniter.MarshalToString(v)
	if err != nil {
		return ""
	}
	return s
}

package utils

import (
	"encoding/json"
	"github.com/json-iterator/go"
	"io"
)

func UnmarshalJSON(data io.Reader, dest interface{}) error {
	err:= json.NewDecoder(data).Decode(dest)
	if err != nil {
		return err
	}
	return nil
}

func MarshalJSON(v interface{}) ([]byte, error) {
	b, err := jsoniter.Marshal(v)
	return b, err
}

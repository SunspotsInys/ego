package jsons

import (
	"bytes"

	jsoniter "github.com/json-iterator/go"
)

type RawMessage = jsoniter.RawMessage

var json jsoniter.API

func init() {
	json = jsoniter.ConfigCompatibleWithStandardLibrary
}

func Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

func Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func JSON2Bytes(v any) *bytes.Reader {
	b, err := json.Marshal(v)
	if err != nil {
		return nil
	}
	return bytes.NewReader(b)
}

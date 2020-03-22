package utils

import "encoding/json"

type jsonConverter struct {
}

func NewJsonConverter() *jsonConverter {
	return &jsonConverter{}
}

func (j jsonConverter) Encode(obj interface{}) ([]byte, error) {
	if obj != nil {
		return json.Marshal(obj)
	}
	return nil, nil
}

func (j jsonConverter) Decode(data []byte, obj interface{}) error {
	return json.Unmarshal(data, obj)
}

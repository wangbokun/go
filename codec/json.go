package codec

import (
	"encoding/json"
 )

// JSON json
type JSON struct{}

// Encode encode
func (JSON) Encode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// Decode decode
func (JSON) Decode(b []byte, v interface{}) error {
	return json.Unmarshal(b, v)
}

// Format format
func (json JSON) Format(dest, src interface{}) error {
 	b, err := json.Encode(src)
	if err != nil {
 		return err
	} 
	return json.Decode(b, dest)
}

func (JSON) String() string {
	return "json"
}

// NewJSONCodec jsoncodec
func NewJSONCodec() JSON {
 	return JSON{}
}

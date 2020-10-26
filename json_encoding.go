package encoding

import (
	"encoding/json"
	"io"
)

var JsonEncoding = NewJsonEncoding()

func NewJsonEncoding() Encoding {
	return &jsonEncoding{}
}

type jsonEncoding struct{}

func (e *jsonEncoding) Marshal(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func (e *jsonEncoding) Unmarshal(b []byte, data interface{}) error {
	if !isPointer(data) {
		return ErrNotAPointer
	}
	return json.Unmarshal(b, data)
}

func (e *jsonEncoding) Encode(data interface{}, w io.Writer) error {
	return json.NewEncoder(w).Encode(data)
}

func (e *jsonEncoding) Decode(r io.Reader, data interface{}) error {
	if !isPointer(data) {
		return ErrNotAPointer
	}
	return json.NewDecoder(r).Decode(data)
}

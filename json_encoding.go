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

func (e *jsonEncoding) Encode(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func (e *jsonEncoding) Decode(b []byte, data interface{}) error {
	if !isPointer(data) {
		return ErrNotAPointer
	}
	return json.Unmarshal(b, data)
}

func (e *jsonEncoding) StreamEncode(data interface{}, w io.Writer) error {
	return json.NewEncoder(w).Encode(data)
}

func (e *jsonEncoding) StreamDecode(r io.Reader, data interface{}) error {
	if !isPointer(data) {
		return ErrNotAPointer
	}
	return json.NewDecoder(r).Decode(data)
}

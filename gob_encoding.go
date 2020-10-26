package encoding

import (
	"encoding/gob"
	"io"
)

var GobEncoding = NewGobEncoding()

func NewGobEncoding() StreamEncoding {
	return &gobEncoding{}
}

type gobEncoding struct{}

func (e *gobEncoding) Encode(data interface{}, w io.Writer) error {
	enc := gob.NewEncoder(w)
	return enc.Encode(data)
}

func (e *gobEncoding) Decode(r io.Reader, data interface{}) error {
	if !isPointer(data) {
		return ErrNotAPointer
	}
	dec := gob.NewDecoder(r)
	return dec.Decode(data)
}

package encoding

import (
	"encoding/hex"
	"io"
)

var HexEncoding = NewHexEncoding()

type hexEncoding struct{}

func (e hexEncoding) NewReader(upstream io.Reader) (io.ReadCloser, error) {
	return io.NopCloser(hex.NewDecoder(upstream)), nil
}

func (e hexEncoding) NewWriter(downstream io.Writer) (io.WriteCloser, error) {
	return nopWriterCloser{hex.NewEncoder(downstream)}, nil
}

func NewHexEncoding() ByteEncoding {
	return NewReadWriteEncoding(&hexEncoding{})
}

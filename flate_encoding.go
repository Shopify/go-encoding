package encoding

import (
	"compress/flate"
	"io"
)

var FlateEncoding = NewFlateEncoding(flate.DefaultCompression, nil)

type flateEncoding struct {
	level int
	dict  []byte
}

func (e flateEncoding) NewReader(upstream io.Reader) (io.ReadCloser, error) {
	return flate.NewReaderDict(upstream, e.dict), nil
}

func (e flateEncoding) NewWriter(downstream io.Writer) (io.WriteCloser, error) {
	return flate.NewWriterDict(downstream, e.level, e.dict)
}

func NewFlateEncoding(level int, dict []byte) ByteEncoding {
	return NewReadWriteEncoding(&flateEncoding{level: level, dict: dict})
}

package encoding

import (
	"compress/zlib"
	"io"
)

var ZlibEncoding = NewZlibEncoding(zlib.DefaultCompression, nil)

type zlibEncoding struct {
	level int
	dict  []byte
}

func (e zlibEncoding) NewReader(upstream io.Reader) (io.ReadCloser, error) {
	return zlib.NewReaderDict(upstream, e.dict)
}

func (e zlibEncoding) NewWriter(downstream io.Writer) (io.WriteCloser, error) {
	return zlib.NewWriterLevelDict(downstream, e.level, e.dict)
}

func NewZlibEncoding(level int, dict []byte) ByteEncoding {
	return NewReadWriteEncoding(&zlibEncoding{level: level, dict: dict})
}

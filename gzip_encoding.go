package encoding

import (
	"compress/gzip"
	"io"
)

var GzipEncoding = NewGzipEncoding(gzip.DefaultCompression)

type gzipEncoding struct {
	level int
}

func (e gzipEncoding) NewReader(upstream io.Reader) (io.ReadCloser, error) {
	return gzip.NewReader(upstream)
}

func (e gzipEncoding) NewWriter(downstream io.Writer) (io.WriteCloser, error) {
	return gzip.NewWriterLevel(downstream, e.level)
}

func NewGzipEncoding(level int) ByteEncoding {
	return NewReadWriteEncoding(&gzipEncoding{level: level})
}

package encoding

import (
	"io"
	"io/ioutil"
)

var NoopEncoding = NewNoopEncoding()

func NewNoopEncoding() ByteEncoding {
	return &noopEncoding{}
}

type noopEncoding struct{}

func (e noopEncoding) StreamEncode(downstream io.Writer) (io.WriteCloser, error) {
	return nopWriterCloser{downstream}, nil
}

func (e noopEncoding) StreamDecode(upstream io.Reader) (io.ReadCloser, error) {
	return ioutil.NopCloser(upstream), nil
}

func (e noopEncoding) Encode(src []byte) ([]byte, error) {
	return src, nil
}

func (e noopEncoding) Decode(src []byte) ([]byte, error) {
	return src, nil
}

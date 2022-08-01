package encoding

import (
	"bytes"
	"io"
	"io/ioutil"
)

type ReaderWriterBuilder interface {
	NewReader(io.Reader) (io.ReadCloser, error)
	NewWriter(io.Writer) (io.WriteCloser, error)
}

func NewReadWriteEncoding(builder ReaderWriterBuilder) ByteEncoding {
	return &readWriteEncoding{
		builder: builder,
	}
}

type readWriteEncoding struct {
	builder ReaderWriterBuilder
}

func (e readWriteEncoding) StreamEncode(downstream io.Writer) (io.WriteCloser, error) {
	return e.builder.NewWriter(downstream)
}

func (e readWriteEncoding) StreamDecode(upstream io.Reader) (io.ReadCloser, error) {
	return e.builder.NewReader(upstream)
}

func (e readWriteEncoding) Encode(src []byte) ([]byte, error) {
	var buf bytes.Buffer
	w, err := e.builder.NewWriter(&buf)
	if err != nil {
		return nil, err
	}

	if _, err := w.Write(src); err != nil {
		return nil, err
	}

	if err := w.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (e readWriteEncoding) Decode(src []byte) ([]byte, error) {
	r, err := e.builder.NewReader(bytes.NewReader(src))
	if err != nil {
		return nil, err
	}
	return readAllClose(r)
}

func readAllClose(r io.ReadCloser) ([]byte, error) {
	dst, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	if err := r.Close(); err != nil {
		return nil, err
	}

	return dst, nil
}

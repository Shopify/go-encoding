package encoding

import (
	"bytes"
	"io"
	"io/ioutil"
)

// NewValueEncoding creates an Encoding respecting the ValueEncoding interface from a StreamEncoding
func NewValueEncoding(e StreamEncoding) Encoding {
	return &valueEncoding{e}
}

type valueEncoding struct {
	StreamEncoding
}

func (e *valueEncoding) Encode(data interface{}) ([]byte, error) {
	var encoded bytes.Buffer
	if err := e.StreamEncode(data, &encoded); err != nil {
		return nil, err
	}

	return encoded.Bytes(), nil
}

func (e *valueEncoding) Decode(b []byte, data interface{}) error {
	reader := bytes.NewReader(b)
	return e.StreamDecode(reader, data)
}

// NewStreamEncoding creates an Encoding respecting the StreamEncoding interface from a ValueEncoding
func NewStreamEncoding(e ValueEncoding) Encoding {
	return &streamEncoding{e}
}

type streamEncoding struct {
	ValueEncoding
}

func (b *streamEncoding) StreamEncode(data interface{}, w io.Writer) error {
	encoded, err := b.Encode(data)
	if err != nil {
		return err
	}
	_, err = w.Write(encoded)
	return err
}

func (b *streamEncoding) StreamDecode(r io.Reader, data interface{}) error {
	decoded, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return b.Decode(decoded, data)
}

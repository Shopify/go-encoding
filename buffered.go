package encoding

import (
	"bytes"
	"io"
	"io/ioutil"
)

// NewBufferedMarshalling creates an Encoding respecting the MarshalEncoding interface from a StreamEncoding
func NewBufferedMarshalling(e StreamEncoding) Encoding {
	return &bufferedMarshalling{e}
}

type bufferedMarshalling struct {
	StreamEncoding
}

func (e *bufferedMarshalling) Marshal(data interface{}) ([]byte, error) {
	var encoded bytes.Buffer
	if err := e.Encode(data, &encoded); err != nil {
		return nil, err
	}

	return encoded.Bytes(), nil
}

func (e *bufferedMarshalling) Unmarshal(b []byte, data interface{}) error {
	reader := bytes.NewReader(b)
	return e.Decode(reader, data)
}

// NewBufferedEncoding creates an Encoding respecting the StreamEncoding interface from a MarshalEncoding
func NewBufferedEncoding(e MarshalEncoding) Encoding {
	return &bufferedEncoding{e}
}

type bufferedEncoding struct {
	MarshalEncoding
}

func (b *bufferedEncoding) Encode(data interface{}, w io.Writer) error {
	encoded, err := b.Marshal(data)
	if err != nil {
		return err
	}
	_, err = w.Write(encoded)
	return err
}

func (b *bufferedEncoding) Decode(r io.Reader, data interface{}) error {
	decoded, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return b.Unmarshal(decoded, data)
}

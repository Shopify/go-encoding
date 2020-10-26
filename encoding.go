package encoding

import "io"

type Encoding interface {
	MarshalEncoding
	StreamEncoding
}

type MarshalEncoding interface {
	Marshaller
	Unmarshaller
}

type Marshaller interface {
	Marshal(data interface{}) ([]byte, error)
}

type Unmarshaller interface {
	Unmarshal(b []byte, data interface{}) error
}

type StreamEncoding interface {
	StreamEncoder
	StreamDecoder
}

type StreamEncoder interface {
	Encode(data interface{}, w io.Writer) error
}

type StreamDecoder interface {
	Decode(r io.Reader, data interface{}) error
}

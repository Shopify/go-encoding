package encoding

import "io"

type Encoding interface {
	ValueEncoding
	StreamEncoding
}

type ValueEncoding interface {
	ValueEncoder
	ValueDecoder
}

type ValueEncoder interface {
	Encode(data interface{}) ([]byte, error)
}

type ValueDecoder interface {
	Decode(b []byte, data interface{}) error
}

type StreamEncoding interface {
	StreamEncoder
	StreamDecoder
}

type StreamEncoder interface {
	StreamEncode(data interface{}, w io.Writer) error
}

type StreamDecoder interface {
	StreamDecode(r io.Reader, data interface{}) error
}

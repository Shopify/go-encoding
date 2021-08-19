package encoding

import "io"

type ByteEncoding interface {
	ByteArrayEncoding
	ByteStreamEncoding
}

type ByteArrayEncoding interface {
	ByteArrayEncoder
	ByteArrayDecoder
}

type ByteArrayEncoder interface {
	Encode([]byte) ([]byte, error)
}

type ByteArrayDecoder interface {
	Decode([]byte) ([]byte, error)
}

type ByteStreamEncoding interface {
	ByteStreamEncoder
	ByteStreamDecoder
}

type ByteStreamEncoder interface {
	StreamEncode(io.Writer) (io.WriteCloser, error)
}

type ByteStreamDecoder interface {
	StreamDecode(io.Reader) (io.ReadCloser, error)
}

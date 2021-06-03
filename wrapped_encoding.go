package encoding

import (
	"io"
)

func NewWrappedValueEncoding(encoding ValueEncoding, byteEncodings ...ByteArrayEncoding) ValueEncoding {
	return &chainValueEncoding{
		encoding:     encoding,
		byteEncoding: NewChainByteArrayEncoding(byteEncodings...),
	}
}

type chainValueEncoding struct {
	encoding     ValueEncoding
	byteEncoding ByteArrayEncoding
}

func (e chainValueEncoding) Encode(data interface{}) ([]byte, error) {
	enc, err := e.encoding.Encode(data)
	if err != nil {
		return nil, err
	}

	return e.byteEncoding.Encode(enc)
}

func (e chainValueEncoding) Decode(enc []byte, data interface{}) error {
	dec, err := e.byteEncoding.Decode(enc)
	if err != nil {
		return err
	}

	return e.encoding.Decode(dec, data)

}

func NewWrappedStreamEncoding(encoding StreamEncoding, byteEncodings ...ByteStreamEncoding) StreamEncoding {
	return &chainStreamEncoding{
		encoding:     encoding,
		byteEncoding: NewChainByteStreamEncoding(byteEncodings...),
	}
}

type chainStreamEncoding struct {
	encoding     StreamEncoding
	byteEncoding ByteStreamEncoding
}

func (e chainStreamEncoding) StreamEncode(data interface{}, w io.Writer) error {
	bw, err := e.byteEncoding.StreamEncode(w)
	if err != nil {
		return err
	}

	if err := e.encoding.StreamEncode(data, bw); err != nil {
		return err
	}

	if err := bw.Close(); err != nil {
		return err
	}

	return nil
}

func (e chainStreamEncoding) StreamDecode(r io.Reader, data interface{}) error {
	br, err := e.byteEncoding.StreamDecode(r)
	if err != nil {
		return err
	}

	if err := e.encoding.StreamDecode(br, data); err != nil {
		return err
	}

	if err := br.Close(); err != nil {
		return err
	}

	return nil
}

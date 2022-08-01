package encoding

import (
	"encoding/base32"
	"io"
	"io/ioutil"
)

var (
	Base32StdEncoding = NewBase32Encoding(base32.StdEncoding)
	Base32URLEncoding = NewBase32Encoding(base32.HexEncoding)
)

func NewBase32Encoding(encoding *base32.Encoding) ByteEncoding {
	return &base32Encoding{encoding: encoding}
}

type base32Encoding struct {
	encoding *base32.Encoding
}

func (e base32Encoding) StreamEncode(downstream io.Writer) (io.WriteCloser, error) {
	w := base32.NewEncoder(e.encoding, downstream)
	return writeCloser{w, downstream}, nil
}

func (e base32Encoding) StreamDecode(upstream io.Reader) (io.ReadCloser, error) {
	return ioutil.NopCloser(base32.NewDecoder(e.encoding, upstream)), nil
}

func (e base32Encoding) Encode(src []byte) ([]byte, error) {
	dst := make([]byte, e.encoding.EncodedLen(len(src)))
	e.encoding.Encode(dst, src)
	return dst, nil
}

func (e base32Encoding) Decode(src []byte) ([]byte, error) {
	dst := make([]byte, e.encoding.DecodedLen(len(src)))
	n, err := e.encoding.Decode(dst, src)
	return dst[:n], err
}

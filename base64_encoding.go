package encoding

import (
	"encoding/base64"
	"io"
	"io/ioutil"
)

var (
	Base64StdEncoding = NewBase64Encoding(base64.StdEncoding)
	Base64URLEncoding = NewBase64Encoding(base64.URLEncoding)
)

func NewBase64Encoding(encoding *base64.Encoding) ByteEncoding {
	return &base64Encoding{encoding: encoding}
}

type base64Encoding struct {
	encoding *base64.Encoding
}

func (e base64Encoding) StreamEncode(downstream io.Writer) (io.WriteCloser, error) {
	return base64.NewEncoder(e.encoding, downstream), nil
}

func (e base64Encoding) StreamDecode(upstream io.Reader) (io.ReadCloser, error) {
	return ioutil.NopCloser(base64.NewDecoder(e.encoding, upstream)), nil
}

func (e base64Encoding) Encode(src []byte) ([]byte, error) {
	dst := make([]byte, e.encoding.EncodedLen(len(src)))
	e.encoding.Encode(dst, src)
	return dst, nil
}

func (e base64Encoding) Decode(src []byte) ([]byte, error) {
	dst := make([]byte, e.encoding.DecodedLen(len(src)))
	n, err := e.encoding.Decode(dst, src)
	return dst[:n], err
}

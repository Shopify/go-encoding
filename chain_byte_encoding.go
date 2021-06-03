package encoding

import (
	"io"
)

func NewChainByteArrayEncoding(encodings ...ByteArrayEncoding) ByteArrayEncoding {
	switch len(encodings) {
	case 0:
		return NoopEncoding
	case 1:
		return encodings[0]
	default:
		return chainByteArrayEncoding(encodings)
	}
}

type chainByteArrayEncoding []ByteArrayEncoding

func (e chainByteArrayEncoding) Encode(b []byte) ([]byte, error) {
	var err error
	for _, layer := range e {
		b, err = layer.Encode(b)
		if err != nil {
			return nil, err
		}
	}

	return b, nil
}

func (e chainByteArrayEncoding) Decode(b []byte) ([]byte, error) {
	var err error
	for i := len(e) - 1; i >= 0; i-- {
		b, err = e[i].Decode(b)
		if err != nil {
			return nil, err
		}
	}
	return b, err
}

func NewChainByteStreamEncoding(encodings ...ByteStreamEncoding) ByteStreamEncoding {
	switch len(encodings) {
	case 0:
		return NoopEncoding
	case 1:
		return encodings[0]
	default:
		return chainByteStreamEncoding(encodings)
	}
}

type chainByteStreamEncoding []ByteStreamEncoding

func (e chainByteStreamEncoding) StreamEncode(w io.Writer) (io.WriteCloser, error) {
	wc := make(writeCloser, len(e)+1)
	wc[len(e)] = w

	var err error
	for i := len(e) - 1; i >= 0; i-- {
		w, err = e[i].StreamEncode(w)
		if err != nil {
			return nil, err
		}
		wc[i] = w
	}

	return wc, nil
}

func (e chainByteStreamEncoding) StreamDecode(r io.Reader) (io.ReadCloser, error) {
	rc := make(readCloser, len(e)+1)
	rc[len(e)] = r

	var err error
	for i := len(e) - 1; i >= 0; i-- {
		r, err = e[i].StreamDecode(r)
		if err != nil {
			return nil, err
		}
		rc[i] = r
	}
	return rc, nil
}

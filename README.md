# Go Encoding

go-encoding is a Go package which abstracts various encoding mechanisms under a unified API.

# Requirements

- [Go 1.13+](http://golang.org/dl/)

# Installation

```console
$ go get github.com/Shopify/go-encoding
```

# Usage

Encodings in this package follow either or both of these [interfaces](encoding.go):

```go
type ValueEncoding interface {
    Encode(data interface{}) ([]byte, error)
    Decode(b []byte, data interface{}) error
}

type StreamEncoding interface {
    StreamEncode(data interface{}, w io.Writer) error
    StreamDecode(r io.Reader, data interface{}) error
}

type ByteArrayEncoding interface {
    Encode([]byte) ([]byte, error)
    Decode([]byte) ([]byte, error)
}

type ByteStreamEncoding interface {
    StreamEncode(io.Writer) (io.WriteCloser, error)
    StreamDecode(io.Reader) (io.ReadCloser, error)
}
```

See [`example_test.go`](example_test.go) for how it all fits together.

## Encodings

Used to encode arbitrary Go types to a byte array.

|   Name  | ValueEncoding | StreamEncoding |                      Implementation                      |
|:--------|:-------------:|:--------------:|:---------------------------------------------------------|
| Literal |       ✔       |                | [`strconv`](https://golang.org/pkg/strconv/)             |
| Json    |       ✔       |        ✔       | [`encoding/json`](https://golang.org/pkg/encoding/json/) |
| Gob     |               |        ✔       | [`encoding/gob`](https://golang.org/pkg/encoding/gob/)   |

To use a `ValueEncoding` as a `StreamEncoding`, you can wrap with `NewStreamEncoding`.

Or vice-versa by using `NewValueEncoding`.

## Byte Encodings

Used to encode byte arrays to byte arrays, which is useful to do processing like compression, encryption, or base64.

|   Name  | ByteArrayEncoding | ByteStreamEncoding |                        Implementation                        |
|:--------|:-----------------:|:------------------:|:-------------------------------------------------------------|
| Base32  |         ✔         |          ✔         | [`encoding/base32`](https://golang.org/pkg/encoding/base32/) |
| Base64  |         ✔         |          ✔         | [`encoding/base64`](https://golang.org/pkg/encoding/base64/) |
| Flate   |         ✔         |          ✔         | [`compress/flate`](https://golang.org/pkg/compress/flate/)   |
| Gzip    |         ✔         |          ✔         | [`compress/gzip`](https://golang.org/pkg/compress/gzip/)     |
| Hex     |         ✔         |          ✔         | [`encoding/hex`](https://golang.org/pkg/encoding/hex/)       |
| Noop    |         ✔         |          ✔         | _passthrough_                                                |
| Zlib    |         ✔         |          ✔         | [`compress/zlib`](https://golang.org/pkg/compress/zlib/)     |

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
```

## Encodings

|   Name  | ValueEncoding | StreamEncoding |                      Implementation                      |
|:--------|:-------------:|:--------------:|:---------------------------------------------------------|
| Literal |       ✔       |                | [`strconv`](https://golang.org/pkg/strconv/)             |
| Json    |       ✔       |        ✔       | [`encoding/json`](https://golang.org/pkg/encoding/json/) |
| Gob     |               |        ✔       | [`encoding/gob`](https://golang.org/pkg/encoding/gob/)   |

To use a `ValueEncoding` as a `StreamEncoding`, you can wrap with `NewStreamEncoding`.

Or vice-versa by using `NewValueEncoding`.

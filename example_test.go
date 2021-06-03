package encoding_test

import (
	"fmt"

	"github.com/Shopify/go-encoding"
)

func ExampleValueEncoding() {
	e := encoding.NewWrappedValueEncoding(encoding.JSONEncoding, encoding.GzipEncoding, encoding.Base64StdEncoding)

	enc, _ := e.Encode(123)
	var num int
	_ = e.Decode(enc, &num)
	fmt.Printf("literal number, encoded: %s, decoded: %d\n", enc, num)

	enc, _ = e.Encode("123")
	var s string
	_ = e.Decode(enc, &s)
	fmt.Printf("literal string, encoded: %s, decoded: %s\n", enc, s)

	enc, _ = e.Encode(map[string]string{"foo": "bar"})
	var m map[string]string
	_ = e.Decode(enc, &m)
	fmt.Printf("json map, encoded: %s, decoded: %+v\n", enc, m)

	// Output:
	// literal number, encoded: H4sIAAAAAAAA/zI0MgYEAAD//9JjSIgDAAAA, decoded: 123
	// literal string, encoded: H4sIAAAAAAAA/1IyNDJWAgQAAP//lwFQMwUAAAA=, decoded: 123
	// json map, encoded: H4sIAAAAAAAA/6pWSsvPV7JSSkosUqoFBAAA///v9Sv+DQAAAA==, decoded: map[foo:bar]
}

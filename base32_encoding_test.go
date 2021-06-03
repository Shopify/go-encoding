package encoding

import (
	"encoding/base32"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestBase32Encoding(t *testing.T) {
	encodings := map[string]*base32.Encoding{
		"std": base32.StdEncoding,
		"hex": base32.HexEncoding,
	}
	for name, encoding := range encodings {
		t.Run(name, func(t *testing.T) {
			suite.Run(t, NewByteEncodingSuite(NewBase32Encoding(encoding)))
		})
	}
}

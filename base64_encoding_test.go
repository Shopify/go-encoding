package encoding

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestBase64Encoding(t *testing.T) {
	encodings := map[string]*base64.Encoding{
		"std":    base64.StdEncoding,
		"rawstd": base64.RawStdEncoding,
		"url":    base64.URLEncoding,
		"rawurl": base64.RawURLEncoding,
	}
	for name, encoding := range encodings {
		t.Run(name, func(t *testing.T) {
			suite.Run(t, NewByteEncodingSuite(NewBase64Encoding(encoding)))
		})
	}
}

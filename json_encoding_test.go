package encoding

import (
	"testing"
)

func TestJsonEncoding(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		testBasicEncoding(t, JsonEncoding)
	})

	t.Run("arbitrary", func(t *testing.T) {
		testArbitraryEncoding(t, JsonEncoding)
	})
}

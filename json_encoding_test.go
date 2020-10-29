package encoding

import (
	"testing"
)

func TestJSONEncoding(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		testBasicEncoding(t, JSONEncoding)
	})

	t.Run("arbitrary", func(t *testing.T) {
		testArbitraryEncoding(t, JSONEncoding)
	})
}

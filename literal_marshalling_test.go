package encoding

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLiteralEncoding(t *testing.T) {
	for _, e := range []ValueEncoding{nil, JSONEncoding, NewValueEncoding(GobEncoding)} {
		t.Run(fmt.Sprintf("%T", e), func(t *testing.T) {
			l := NewLiteralEncoding(e)

			t.Run("basic", func(t *testing.T) {
				testBasicEncoding(t, l)
			})

			if e != nil {
				t.Run("arbitrary", func(t *testing.T) {
					testArbitraryEncoding(t, l)
				})
			}
		})
	}

	t.Run("unsupported", func(t *testing.T) {
		enc, err := LiteralEncoding.Encode(struct{}{})
		require.Zero(t, enc)
		require.EqualError(t, err, "not implemented for type struct")
	})
}

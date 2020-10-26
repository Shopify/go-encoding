package encoding

import (
	"encoding/gob"
	"testing"
	"time"
)

func TestGobEncoding(t *testing.T) {
	e := NewBufferedMarshalling(GobEncoding)

	t.Run("basic", func(t *testing.T) {
		testBasicEncoding(t, e)
	})

	t.Run("arbitrary", func(t *testing.T) {
		gob.Register(struct{}{})
		gob.Register(testStruct{})
		gob.Register(time.Time{})
		testArbitraryEncoding(t, e)
	})
}

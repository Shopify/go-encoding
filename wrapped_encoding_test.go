package encoding

import (
	"testing"
)

func TestWrappedValueEncoding(t *testing.T) {
	encoding := NewWrappedValueEncoding(JSONEncoding, NewChainByteArrayEncoding(GzipEncoding, Base64StdEncoding))
	testBasicEncoding(t, encoding)
}

func TestWrappedStreamEncoding(t *testing.T) {
	encoding := NewWrappedStreamEncoding(JSONEncoding, NewChainByteStreamEncoding(GzipEncoding, Base64StdEncoding))
	testBasicEncoding(t, NewValueEncoding(encoding))
}

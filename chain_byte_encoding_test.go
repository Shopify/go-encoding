package encoding

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestChainByteEncoding(t *testing.T) {
	suite.Run(t, &ByteEncodingSuite{
		byteArrayEncoding:  NewChainByteArrayEncoding(GzipEncoding, Base64URLEncoding),
		byteStreamEncoding: NewChainByteStreamEncoding(GzipEncoding, Base64URLEncoding),
	})
}

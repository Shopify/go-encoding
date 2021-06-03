package encoding

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestZlibEncoding(t *testing.T) {
	suite.Run(t, NewByteEncodingSuite(ZlibEncoding))
}

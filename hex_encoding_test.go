package encoding

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestHexEncoding(t *testing.T) {
	suite.Run(t, NewByteEncodingSuite(HexEncoding))
}

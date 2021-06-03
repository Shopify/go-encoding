package encoding

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestFlateEncoding(t *testing.T) {
	suite.Run(t, NewByteEncodingSuite(FlateEncoding))
}

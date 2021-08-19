package encoding

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestNoopEncoding(t *testing.T) {
	suite.Run(t, NewByteEncodingSuite(NoopEncoding))
}

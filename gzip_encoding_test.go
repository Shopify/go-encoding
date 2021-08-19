package encoding

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestGzipEncoding(t *testing.T) {
	suite.Run(t, NewByteEncodingSuite(GzipEncoding))
}

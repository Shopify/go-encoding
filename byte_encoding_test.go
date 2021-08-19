package encoding

import (
	"bytes"

	"github.com/stretchr/testify/suite"
)

func NewByteEncodingSuite(e ByteEncoding) *ByteEncodingSuite {
	return &ByteEncodingSuite{
		byteArrayEncoding:  e,
		byteStreamEncoding: e,
	}
}

type ByteEncodingSuite struct {
	suite.Suite
	byteArrayEncoding  ByteArrayEncoding
	byteStreamEncoding ByteStreamEncoding
}

func (suite ByteEncodingSuite) TestByteArray() {
	if suite.byteArrayEncoding == nil {
		suite.T().Skip()
	}

	tests := [][]byte{
		{},
		{0},
		{0, 0},
		[]byte("123"),
		[]byte(`"123"`),
	}

	for _, input := range tests {
		suite.Run(string(input), func() {
			enc, err := suite.byteArrayEncoding.Encode(input)
			suite.Require().NoError(err)

			dec, err := suite.byteArrayEncoding.Decode(enc)
			suite.Require().NoError(err)

			suite.Require().Equal(input, dec)
		})
	}
}

func (suite ByteEncodingSuite) TestByteStream() {
	if suite.byteStreamEncoding == nil {
		suite.T().Skip()
	}

	tests := [][]byte{
		{},
		{0},
		{0, 0},
		[]byte("123"),
		[]byte(`"123"`),
	}

	for _, input := range tests {
		suite.Run(string(input), func() {
			var buf bytes.Buffer

			w, err := suite.byteStreamEncoding.StreamEncode(&buf)
			suite.Require().NoError(err)

			_, err = w.Write(input)
			suite.Require().NoError(err)

			err = w.Close()
			suite.Require().NoError(err)

			r, err := suite.byteStreamEncoding.StreamDecode(&buf)
			suite.Require().NoError(err)

			dec, err := readAllClose(r)
			suite.Require().NoError(err)

			suite.Require().Equal(input, dec)
		})
	}
}

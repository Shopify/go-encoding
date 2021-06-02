package encoding

// ByteArrayEncrypter is provided as a common interface for encryption operations
type ByteArrayEncrypter interface {
	Encrypt([]byte) ([]byte, error)
	Decrypt([]byte) ([]byte, error)
}

func NewByteArrayEncrypterEncoding(encrypter ByteArrayEncrypter) ByteArrayEncoding {
	return &encrypterEncoding{encrypter: encrypter}
}

type encrypterEncoding struct {
	encrypter ByteArrayEncrypter
}

func (e encrypterEncoding) Encode(src []byte) ([]byte, error) {
	return e.encrypter.Encrypt(src)
}

func (e encrypterEncoding) Decode(src []byte) ([]byte, error) {
	return e.encrypter.Decrypt(src)
}

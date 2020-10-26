package encoding

import (
	"fmt"
	"reflect"
)

var LiteralEncoding = NewLiteralEncoding(nil)

// NewLiteralEncoding is an encoding that will try its best to store the data as is,
// but fallback on another encoder if not possible.
func NewLiteralEncoding(fallback MarshalEncoding) MarshalEncoding {
	return &literalMarshalling{fallback: fallback}
}

type literalMarshalling struct {
	fallback MarshalEncoding
}

func (e *literalMarshalling) Marshal(data interface{}) ([]byte, error) {
	value := reflect.ValueOf(data)
	if t, ok := translators[value.Kind()]; ok {
		return t.encode(value), nil
	}

	if e.fallback == nil {
		return nil, fmt.Errorf("not implemented for type %s", value.Kind())
	}

	return e.fallback.Marshal(data)
}

func (e *literalMarshalling) Unmarshal(b []byte, data interface{}) error {
	if !isPointer(data) {
		return ErrNotAPointer
	}

	value := reflect.ValueOf(data).Elem()
	if t, ok := translators[value.Kind()]; ok {
		return t.decode(b, value)
	}

	if e.fallback == nil {
		return fmt.Errorf("not implemented for type %s", value.Kind())
	}

	return e.fallback.Unmarshal(b, data)
}

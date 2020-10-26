package encoding

import (
	"reflect"
	"strconv"
)

var translators = map[reflect.Kind]translator{
	reflect.Uint:   uintTranslator{strconv.IntSize},
	reflect.Uint8:  uintTranslator{8},
	reflect.Uint16: uintTranslator{16},
	reflect.Uint32: uintTranslator{32},
	reflect.Uint64: uintTranslator{64},

	reflect.Int:   intTranslator{strconv.IntSize},
	reflect.Int8:  intTranslator{8},
	reflect.Int16: intTranslator{16},
	reflect.Int32: intTranslator{32},
	reflect.Int64: intTranslator{64},

	reflect.Float32: floatTranslator{32},
	reflect.Float64: floatTranslator{64},

	reflect.String: stringTranslator{},

	reflect.Bool: boolTranslator{},
}

type translator interface {
	encode(val reflect.Value) []byte
	decode(data []byte, target reflect.Value) error
}

type uintTranslator struct {
	bitSize int
}

func (uintTranslator) encode(val reflect.Value) []byte {
	return []byte(strconv.FormatUint(val.Uint(), 10))
}

func (t uintTranslator) decode(data []byte, target reflect.Value) error {
	val, err := strconv.ParseUint(string(data), 10, t.bitSize)
	if err != nil {
		return err
	}

	target.SetUint(val)
	return nil
}

type intTranslator struct {
	bitSize int
}

func (intTranslator) encode(val reflect.Value) []byte {
	return []byte(strconv.FormatInt(val.Int(), 10))
}

func (t intTranslator) decode(data []byte, target reflect.Value) error {
	val, err := strconv.ParseInt(string(data), 10, t.bitSize)
	if err != nil {
		return err
	}

	target.SetInt(val)
	return nil
}

type floatTranslator struct {
	bitSize int
}

func (t floatTranslator) encode(val reflect.Value) []byte {
	return []byte(strconv.FormatFloat(val.Float(), 'f', -1, t.bitSize))
}

func (t floatTranslator) decode(data []byte, target reflect.Value) error {
	val, err := strconv.ParseFloat(string(data), t.bitSize)
	if err != nil {
		return err
	}

	target.SetFloat(val)
	return nil
}

type stringTranslator struct{}

func (stringTranslator) encode(val reflect.Value) []byte {
	return []byte(val.String())
}

func (stringTranslator) decode(data []byte, target reflect.Value) error {
	target.SetString(string(data))
	return nil
}

type boolTranslator struct{}

func (boolTranslator) encode(val reflect.Value) []byte {
	return []byte(strconv.FormatBool(val.Bool()))
}

func (boolTranslator) decode(data []byte, target reflect.Value) error {
	val, err := strconv.ParseBool(string(data))
	if err != nil {
		return err
	}

	target.SetBool(val)
	return nil
}

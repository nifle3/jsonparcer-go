package jsonparcer_go

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
)

type tagName string

const (
	JsonTag tagName = "json"
)

type Parser interface {
	HasError() bool
	GetError() []error
	Parse(r io.Reader) func(func(string, any) bool)
}

type Marshaller[K any] struct {
	parser Parser
	tag    tagName
}

func NewJsonMarsheller[K any]() Marshaller[K] {
	return Marshaller[K]{
		parser: nil,
		tag:    JsonTag,
	}
}

func newTestMarsheller[K any]() Marshaller[K] {
	return Marshaller[K]{
		parser: testParser{},
		tag:    JsonTag,
	}
}

type fieldInfo struct {
	kind    reflect.Kind
	idField int
}

func (m Marshaller[K]) UnmarshallingByte(readerBytes []byte) (K, error) {
	r := bytes.NewReader(readerBytes)

	return m.Unmarshalling(r)
}

func (m Marshaller[K]) Unmarshalling(r io.Reader) (K, error) {
	var result K
	valueInfo := reflect.ValueOf(&result)
	reflectType := reflect.TypeOf(result)

	typeInfo := make(map[string]fieldInfo)

	for i := range reflectType.NumField() {
		jsonName, ok := reflectType.Field(i).Tag.Lookup(string(m.tag))
		if !ok {
			continue
		}

		typeInfo[jsonName] = fieldInfo{
			kind:    reflectType.Field(i).Type.Kind(),
			idField: i,
		}
	}

	for key, value := range m.parser.Parse(r) {
		field, ok := typeInfo[key]
		if !ok {
			continue
		}

		if field.kind != reflect.ValueOf(value).Kind() {
			return result, fmt.Errorf("type mismatch got %s expected %s",
				reflect.ValueOf(value).Kind(), field.kind)
		}

		if !valueInfo.Elem().Field(field.idField).CanSet() {
			return result, fmt.Errorf("cant set a %s field",
				key)
		}

		valueInfo.Elem().Field(field.idField).Set(reflect.ValueOf(value))
	}

	return result, nil
}

type testParser struct {
}

func (t testParser) HasError() bool {
	return false
}

func (t testParser) GetError() []error {
	return nil
}

func (t testParser) Parse(r io.Reader) func(func(string, any) bool) {
	return func(yield func(string, any) bool) {
		jsonTest := map[string]interface{}{
			"name":      "qwe",
			"last_name": "Qwe",
			"surname":   "QWE",
			"age":       123,
			"age2":      23,
		}

		for key, value := range jsonTest {
			if !yield(key, value) {
				return
			}
		}
	}
}

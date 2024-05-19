package jsonparcer_go

import (
	"fmt"
	"reflect"
)

type fieldInfo struct {
	kind    reflect.Kind
	idField int
}

func Unmarshalling[K any]() (K, error) {
	var result K
	valueInfo := reflect.ValueOf(&result)
	reflectType := reflect.TypeOf(result)

	typeInfo := make(map[string]fieldInfo)

	for i := range reflectType.NumField() {
		jsonName, ok := reflectType.Field(i).Tag.Lookup("json")
		if !ok {
			continue
		}

		typeInfo[jsonName] = fieldInfo{
			kind:    reflectType.Field(i).Type.Kind(),
			idField: i,
		}
	}

	for key, value := range parseJson() {
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

func parseJson() func(func(string, any) bool) {
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

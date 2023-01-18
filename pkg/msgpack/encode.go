package msgpack

import (
	"reflect"
)

func Marahal(v any) ([]byte, error) {
	buffer := []byte{}

	if v == nil {
		buffer = append(buffer, TypeNil)
	} else {
		va := reflect.ValueOf(v)

		switch va.Kind() {
		case reflect.String:
			strLen := va.Len()
			buffer = append(buffer, byte(TypeFixStr)|byte(strLen))
			buffer = append(buffer, []byte(va.String())...)
		}
	}

	return buffer, nil
}

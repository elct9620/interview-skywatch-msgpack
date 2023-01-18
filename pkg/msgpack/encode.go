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
			buffer = append(buffer, encodeString(va)...)
		case reflect.Map:
			elements := va.Len()
			buffer = append(buffer, byte(TypeFixMap)|byte(elements))
			for _, key := range va.MapKeys() {
				buffer = append(buffer, encodeString(key)...)
				buffer = append(buffer, encodeString(va.MapIndex(key))...)
			}
		}
	}

	return buffer, nil
}

func encodeString(va reflect.Value) (buffer []byte) {
	strLen := va.Len()
	buffer = append(buffer, byte(TypeFixStr)|byte(strLen))
	buffer = append(buffer, []byte(va.String())...)

	return buffer
}

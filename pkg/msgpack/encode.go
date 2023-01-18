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
		buffer = append(buffer, encode(va)...)
	}

	return buffer, nil
}

func encode(va reflect.Value) (buffer []byte) {
	switch va.Kind() {
	case reflect.String:
		buffer = append(buffer, encodeString(va)...)
	case reflect.Map:
		buffer = append(buffer, encodeMap(va)...)
	case reflect.Struct:
		buffer = append(buffer, encodeStruct(va)...)
	}

	return buffer
}

func encodeString(va reflect.Value) (buffer []byte) {
	strLen := va.Len()
	buffer = append(buffer, byte(TypeFixStr)|byte(strLen))
	buffer = append(buffer, []byte(va.String())...)

	return buffer
}

func encodeMap(va reflect.Value) (buffer []byte) {
	numElement := va.Len()
	buffer = append(buffer, byte(TypeFixMap)|byte(numElement))
	for _, key := range va.MapKeys() {
		buffer = append(buffer, encodeString(key)...)

		buffer = append(buffer, encode(va.MapIndex(key))...)
	}

	return buffer
}

func encodeStruct(va reflect.Value) (buffer []byte) {
	numField := va.NumField()
	buffer = append(buffer, byte(TypeFixMap)|byte(numField))

	vt := va.Type()
	for i := 0; i < numField; i++ {
		field := vt.Field(i)
		name, ok := field.Tag.Lookup("msgpack")
		if !ok {
			name = field.Name
		}

		buffer = append(buffer, encodeString(reflect.ValueOf(name))...)
		buffer = append(buffer, encode(va.Field(i))...)
	}

	return buffer
}

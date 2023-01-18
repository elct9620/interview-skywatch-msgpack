package msgpack

import (
	"encoding/json"
	"reflect"
)

func FromJSON(data []byte) (buffer []byte, err error) {
	var payload any
	err = json.Unmarshal(data, &payload)
	if err != nil {
		return nil, err
	}

	return Marahal(payload)
}

func Marahal(v any) (buffer []byte, err error) {
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
	case reflect.Bool:
		buffer = append(buffer, encodeBool(va)...)
	case reflect.Int:
		buffer = append(buffer, encodeInt(va)...)
	case reflect.String:
		buffer = append(buffer, encodeString(va)...)
	case reflect.Slice:
		buffer = append(buffer, encodeSlice(va)...)
	case reflect.Map:
		buffer = append(buffer, encodeMap(va)...)
	case reflect.Struct:
		buffer = append(buffer, encodeStruct(va)...)
	case reflect.Interface:
		buffer = append(buffer, encode(va.Elem())...)
	}

	return buffer
}

func encodeBool(va reflect.Value) []byte {
	if va.Bool() == true {
		return []byte{TypeTrue}
	}

	return []byte{TypeFalse}
}

func encodeInt(va reflect.Value) (buffer []byte) {
	value := va.Int()
	if value > 0 {
		buffer = append(buffer, byte(TypePositiveFixInt)|byte(value))
	} else {
		buffer = append(buffer, byte(TypeNegativeFixInt)|byte(value))
	}
	return buffer
}

func encodeString(va reflect.Value) (buffer []byte) {
	strLen := va.Len()
	buffer = append(buffer, byte(TypeFixStr)|byte(strLen))
	buffer = append(buffer, []byte(va.String())...)

	return buffer
}

func encodeSlice(va reflect.Value) (buffer []byte) {
	numElement := va.Len()
	buffer = append(buffer, byte(TypeFixArray)|byte(numElement))

	for i := 0; i < numElement; i++ {
		buffer = append(buffer, encode(va.Index(i))...)
	}

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

package msgpack

import (
	"encoding/binary"
	"encoding/json"
	"math"
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
	case reflect.Float32, reflect.Float64:
		buffer = append(buffer, encodeFloat(va)...)
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
		buffer = append(buffer, TypePositiveFixInt|byte(value))
	} else {
		buffer = append(buffer, TypeNegativeFixInt|byte(value))
	}
	return buffer
}

func encodeFloat(va reflect.Value) (buffer []byte) {
	if va.Kind() == reflect.Float32 {
		buffer = append(buffer, TypeFloat32)
	} else {
		buffer = append(buffer, TypeFloat64)
	}

	var floatBytes []byte = make([]byte, 8)
	binary.BigEndian.PutUint64(floatBytes, math.Float64bits(va.Float()))
	buffer = append(buffer, floatBytes...)

	return buffer
}

func encodeString(va reflect.Value) (buffer []byte) {
	switch strLen := va.Len(); {
	case strLen < FixStrMaxLen:
		buffer = append(buffer, TypeFixStr|byte(strLen))
	case strLen < Str8MaxLen:
		buffer = append(buffer, TypeStr8, byte(strLen))
	case strLen < Str16MaxLen:
		buffer = append(buffer, TypeStr16)
		buffer = append(buffer, toInt16Bytes(uint16(strLen))...)
	case strLen < Str32MaxLen:
		buffer = append(buffer, TypeStr32)
		buffer = append(buffer, toInt32Bytes(uint32(strLen))...)
	}
	buffer = append(buffer, []byte(va.String())...)

	return buffer
}

func encodeSlice(va reflect.Value) (buffer []byte) {
	numElement := va.Len()
	buffer = append(buffer, TypeFixArray|byte(numElement))

	for i := 0; i < numElement; i++ {
		buffer = append(buffer, encode(va.Index(i))...)
	}

	return buffer
}

func encodeMap(va reflect.Value) (buffer []byte) {
	numElement := va.Len()
	buffer = append(buffer, TypeFixMap|byte(numElement))
	for _, key := range va.MapKeys() {
		buffer = append(buffer, encodeString(key)...)

		buffer = append(buffer, encode(va.MapIndex(key))...)
	}

	return buffer
}

func encodeStruct(va reflect.Value) (buffer []byte) {
	numField := va.NumField()
	buffer = append(buffer, TypeFixMap|byte(numField))

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

func toInt16Bytes(v uint16) []byte {
	buffer := make([]byte, 4)
	binary.BigEndian.PutUint16(buffer, v)
	return buffer
}

func toInt32Bytes(v uint32) []byte {
	buffer := make([]byte, 4)
	binary.BigEndian.PutUint32(buffer, v)
	return buffer
}

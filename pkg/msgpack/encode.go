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
	case reflect.Uint:
		buffer = append(buffer, encodeUint(va)...)
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

func encodeUint(va reflect.Value) (buffer []byte) {
	value := va.Int()
	switch {
	case value < PositiveFixIntMax:
		buffer = append(buffer, TypePositiveFixInt|byte(value))
	case value < Uint8Max:
		buffer = append(buffer, TypeUint8, byte(value))
	case value < Uint16Max:
		buffer = append(buffer, TypeUint16)
		buffer = append(buffer, toInt16Bytes(uint16(value))...)
	case value < Uint32Max:
		buffer = append(buffer, TypeUint32)
		buffer = append(buffer, toInt32Bytes(uint32(value))...)
	default:
		buffer = append(buffer, TypeUint64)
		buffer = append(buffer, toInt64Bytes(uint64(value))...)
	}

	return buffer
}

func encodeInt(va reflect.Value) (buffer []byte) {
	value := va.Int()
	if value >= 0 {
		return encodeUint(va)
	}

	switch {
	case value >= NegativeFixIntMin:
		buffer = append(buffer, TypeNegativeFixInt|byte(value))
	case value > Int8Min:
		buffer = append(buffer, TypeInt8)
		buffer = append(buffer, byte(value))
	case value > Int16Min:
		buffer = append(buffer, TypeInt16)
		buffer = append(buffer, toInt16Bytes(uint16(value))...)
	case value > Int32Min:
		buffer = append(buffer, TypeInt32)
		buffer = append(buffer, toInt32Bytes(uint32(value))...)
	default:
		buffer = append(buffer, TypeInt64)
		buffer = append(buffer, toInt64Bytes(uint64(value))...)
	}

	return buffer
}

func encodeFloat(va reflect.Value) (buffer []byte) {
	if va.Kind() == reflect.Float32 {
		buffer = append(buffer, TypeFloat32)
		buffer = append(buffer, toInt32Bytes(math.Float32bits(float32(va.Float())))...)
	} else {
		buffer = append(buffer, TypeFloat64)
		buffer = append(buffer, toInt64Bytes(math.Float64bits(va.Float()))...)
	}

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
	switch {
	case numElement < FixArrayMaxElement:
		buffer = append(buffer, TypeFixArray|byte(numElement))
	case numElement < Array16MaxElement:
		buffer = append(buffer, TypeArray16)
		buffer = append(buffer, toInt16Bytes(uint16(numElement))...)
	case numElement < Array32MaxElement:
		buffer = append(buffer, TypeArray32)
		buffer = append(buffer, toInt32Bytes(uint32(numElement))...)
	}

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
	buffer := make([]byte, 2)
	binary.BigEndian.PutUint16(buffer, v)
	return buffer
}

func toInt32Bytes(v uint32) []byte {
	buffer := make([]byte, 4)
	binary.BigEndian.PutUint32(buffer, v)
	return buffer
}

func toInt64Bytes(v uint64) []byte {
	buffer := make([]byte, 8)
	binary.BigEndian.PutUint64(buffer, v)
	return buffer
}

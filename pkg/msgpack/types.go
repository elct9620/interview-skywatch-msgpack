package msgpack

const (
	TypeNil            byte = 0xc0
	TypeFalse               = 0xc2
	TypeTrue                = 0xc3
	TypePositiveFixInt      = 0b00000000
	TypeNegativeFixInt      = 0b11100000
	TypeUint8               = 0xcc
	TypeUint16              = 0xcd
	TypeUint32              = 0xce
	TypeUint64              = 0xcf
	TypeInt8                = 0xd0
	TypeInt16               = 0xd1
	TypeInt32               = 0xd2
	TypeInt64               = 0xd3
	TypeFloat32             = 0xca
	TypeFloat64             = 0xcb
	TypeFixStr              = 0b10100000
	TypeStr8                = 0xd9
	TypeStr16               = 0xda
	TypeStr32               = 0xdb
	TypeFixArray            = 0b10010000
	TypeArray16             = 0xdc
	TypeArray32             = 0xdd
	TypeFixMap              = 0b10000000
	TypeMap16               = 0xde
	TypeMap32               = 0xdf
)

const (
	PositiveFixIntMax int64 = 127
	Uint8Max                = 2 << 7
	Uint16Max               = 2 << 15
	Uint32Max               = 2 << 31
	NegativeFixIntMin       = -32
	Int8Min                 = -(2 << 7) / 2
	Int16Min                = -(2 << 15) / 2
	Int32Min                = -(2 << 31) / 2
)

const (
	FixStrMaxLen = 32
	Str8MaxLen   = 2 << 7
	Str16MaxLen  = 2 << 15
	Str32MaxLen  = 2 << 31
)

const (
	FixArrayMaxElement = 15
	Array16MaxElement  = 2 << 15
	Array32MaxElement  = 2 << 31
)

const (
	FixMapMaxElement = 15
	Map16MaxElement  = 2 << 15
	Map32MaxElement  = 2 << 31
)

package msgpack

const (
	TypeNil            byte = 0xc0
	TypeFalse               = 0xc2
	TypeTrue                = 0xc3
	TypePositiveFixInt      = 0b00000000
	TypeNegativeFixInt      = 0b11100000
	TypeFloat32             = 0xca
	TypeFloat64             = 0xcb
	TypeFixStr              = 0b10100000
	TypeStr8                = 0xd9
	TypeStr16               = 0xda
	TypeStr32               = 0xdb
	TypeFixArray            = 0b10010000
	TypeFixMap              = 0b10000000
)

const (
	FixStrMaxLen = 32
	Str8MaxLen   = 2 << 7
	Str16MaxLen  = 2 << 15
	Str32MaxLen  = 2 << 31
)

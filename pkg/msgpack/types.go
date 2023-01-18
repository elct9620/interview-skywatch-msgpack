package msgpack

const (
	TypeNil            = 0xc0
	TypeFalse          = 0xc2
	TypeTrue           = 0xc3
	TypePositiveFixInt = 0b00000000
	TypeNegativeFixInt = 0b11100000
	TypeFloat32        = 0xca
	TypeFloat64        = 0xcb
	TypeFixStr         = 0b10100000
	TypeFixArray       = 0b10010000
	TypeFixMap         = 0b10000000
)

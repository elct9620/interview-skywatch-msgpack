package msgpack

func Marahal(v any) ([]byte, error) {
	buffer := []byte{}

	switch v.(type) {
	case nil:
		buffer = append(buffer, NilType)
	}

	return buffer, nil
}

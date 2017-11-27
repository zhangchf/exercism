package variablelengthquantity

import (
	"errors"
)

func EncodeVarint(input []uint32) (output []byte) {
	output = make([]byte, 0)
	sign := uint32(1 << 7)
	for _, i := range input {
		encoded := []byte{byte(i % sign)}
		for i >>= 7; i > 0; i >>= 7 {
			encoded = append([]byte{byte(i%sign + sign)}, encoded...)
		}
		output = append(output, encoded...)
	}
	return output
}

func DecodeVarint(input []byte) (output []uint32, err error) {
	var num uint32
	incomplete := false
	for _, b := range input {
		num <<= 7
		num += uint32(b & 0x7f)
		incomplete = (b&0x80 != 0)
		if !incomplete {
			output = append(output, num)
			num = 0
		}
	}
	if incomplete {
		return nil, errors.New("Invalid input")
	}
	return
}

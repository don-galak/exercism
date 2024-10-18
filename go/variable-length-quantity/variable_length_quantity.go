package variablelengthquantity

import "errors"

func reverse(in []byte) {
	for i, j := 0, len(in)-1; i < j; i, j = i+1, j-1 {
		in[i], in[j] = in[j], in[i]
	}
}

func EncodeVarint(input []uint32) (out []byte) {
	for _, el := range input {
		extracted := []byte{byte(el % 128)}
		for el >>= 7; el != 0; el >>= 7 {
			println(el, el>>7)
			extracted = append(extracted, 128+byte(el%128))
		}
		reverse(extracted)
		out = append(out, extracted...)
	}
	return
}

func DecodeVarint(input []byte) ([]uint32, error) {
	var d uint32
	var complete bool
	decodedInts := make([]uint32, 0)
	for _, b := range input {
		d += uint32(b & 0x7F)
		complete = (b&0x80 == 0)
		if complete {
			decodedInts = append(decodedInts, d)
			d = 0
			continue
		}
		d <<= 7
	}
	if !complete {
		return nil, errors.New("incomplete sequence")
	}
	return decodedInts, nil
}

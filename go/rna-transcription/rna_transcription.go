package strand

import "bytes"

const testVersion = 3

var dna2rnaMap = map[byte]byte{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) (rna string) {
	var buffer = bytes.Buffer{}

	for _, r := range dna {
		buffer.WriteByte(dna2rnaMap[byte(r)])
	}
	return buffer.String()
}

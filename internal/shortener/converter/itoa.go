package converter

import (
	"fmt"
	"runtime"
)

var (
	// letterRunes is list of characters that we use in short URL. It contains numbers, uppercase and lowercase letters
	elems = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
)

const (
	_1stIndex int = 0
	_2ndIndex int = 10
	_3rdIndex int = 20
	_4ndIndex int = 30
)

// ItoA converts uint32 to a string which is based 62 by runes stored in 'elems' variable.
// Expected result is kind of obfuscated.
func ItoA(i uint32) (string, error) {
	switch {
	case i >= 62*62*62*62:
		return "", fmt.Errorf("too large int to convert")
	default:
		res := convertToString(i)
		runtime.GC()
		return res, nil
	}
}

// convertToString converts validated uint32 to a 5 character string
// which is unique for each number.
func convertToString(i uint32) string {
	indexes := make([]uint32, 4)

	for j := 0; j < 3; j++ {
		indexes[j] = i % 62
		i = (i - indexes[j]) / 62
	}

	return fmt.Sprintf(
		"%c%c%c%c",
		elems[(int(indexes[3])+_4ndIndex)%62],
		elems[(int(indexes[2])+_3rdIndex)%62],
		elems[(int(indexes[1])+_2ndIndex)%62],
		elems[(int(indexes[0])+_1stIndex)%62],
	)
}

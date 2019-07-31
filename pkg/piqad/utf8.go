package piqad

import "errors"

// Eng2Piq converts and english symbol to a pIqaD letter name

// ToUTF8 returns an UTF-8 code for a given pIqaD letter name
func ToUTF8(input string) (int, error) {
	code, ok := piqadMap[input]
	if !ok {
		return 0, errors.New("transliterator: unsupported pIqaD letter name")
	}

	return code, nil
}

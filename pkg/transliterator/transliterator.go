package transliterator

import "errors"
import "strings"

// Eng2Piq converts and english symbol to a pIqaD letter name
func Eng2Piq(eng rune) (string, error) {
	downcased := strings.ToLower(string(eng))
	piq, ok := eng2piqadMap[downcased]

	if !ok {
		return "", errors.New("transliterator: unsupported english symbol")
	}

	return piq, nil
}

// Piq2UTF8Code returns an UTF-8 code for a given pIqaD letter name
func Piq2UTF8Code(input string) (int, error) {
	code, ok := piqad2utf8Map[input]
	if !ok {
		return 0, errors.New("transliterator: unsupported pIqaD letter name")
	}

	return code, nil
}

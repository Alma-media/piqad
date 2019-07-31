package piqad

import (
	"fmt"
	"strings"
)

// Transliterate transliterates an english string to piQaD alphabet string
func Transliterate(str string) (string, error) {
	str = strings.ToLower(str)
	buf := ""

	for i, r := range str {
		symb, ok := englishMap[string(r)]
		if !ok {
			return "", fmt.Errorf("piqad.Transliterator: cannot match symbol '%s' at pos %d", string(r), i+1)
		}
		buf += symb
	}

	return buf, nil
}

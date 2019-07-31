package piqad

import (
	"errors"
	"strings"
)

// EnglishTransliterator implements a transliterator from english to klingon alphabet
type EnglishTransliterator struct {
	pos   int
	runes []rune
}

// NewEng2PiqTransliterator returns a new instance of a transliterator
func NewEng2PiqTransliterator(input string) *EnglishTransliterator {
	runes := []rune(input)
	return &EnglishTransliterator{
		runes: runes,
		pos:   0,
	}
}

// Next increases a cursor position and returns true if it's possible to continue iterate
// Pay attention that it's not thread-safe function
func (t *EnglishTransliterator) Next() bool {
	if t.pos >= len(t.runes) {
		return false
	}

	t.pos++
	return true
}

// Get returns a pIqaD letter name matched to an english symbol under a cursor position
func (t *EnglishTransliterator) Get() (string, error) {
	r := t.runes[t.pos]
	return eng2Piq(r)
}

func eng2Piq(eng rune) (string, error) {
	downcased := strings.ToLower(string(eng))
	piq, ok := englishMap[downcased]

	if !ok {
		return "", errors.New("transliterator: unsupported english symbol")
	}

	return piq, nil
}

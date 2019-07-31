package piqad

import (
	"errors"
)

// Tokenizer implements a tokenizer that extracts pIqaD tokens from a string
// Example: "mnng" contains 3 pIqaD letters (i.e. tokens): "m", "n", "ng"
type Tokenizer struct {
	runes        []rune
	token        string
	pos          int
	maxTokenSize int
}

// NewPiqTokenizer buils a tokenizer from an input string
func NewPiqTokenizer(input string) *Tokenizer {
	return &Tokenizer{
		runes:        []rune(input),
		maxTokenSize: 3,
	}
}

// Next changes cursor position to a next token
// Not thread-safe!
func (t *Tokenizer) Next() bool {
	current := ""
	t.token = ""
	found := false
	i := t.pos

	for i < len(t.runes) {
		ch := string(t.runes[i])
		i++
		current += ch

		if isValidToken(current) {
			t.token = current
			t.pos = i
			found = true
		}

		if i-t.pos > t.maxTokenSize {
			break
		}
	}

	return found
}

// Get returns a last found token
func (t *Tokenizer) Get() (string, error) {
	if t.token == "" {
		return "", errors.New("transliterator: cannot find any token")
	}
	return t.token, nil
}

func isValidToken(token string) bool {
	_, ok := piqadMap[token]
	return ok
}

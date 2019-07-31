package piqad

import "errors"

// Tokenizer implements a tokenizer that extracts pIqaD tokens from a string
// Example: "mnng" contains 3 pIqaD letters (i.e. tokens): "m", "n", "ng"
type Tokenizer struct {
	runes []rune // Runes array used to scan over
	token string // Last found token
	done  bool   // Scan has finished
	err   error  // An error
	pos   int    // Current Tokenizer position in runes
}

// NewTokenizer buils a tokenizer from an input string
func NewTokenizer(input string) *Tokenizer {
	return &Tokenizer{
		runes: []rune(input),
	}
}

// Err returns the last error that was encountered by the Tokenizer.
func (t *Tokenizer) Err() error {
	return t.err
}

// Scan changes cursor position to a next token
// Not thread-safe!
func (t *Tokenizer) Scan() bool {
	if t.done {
		return false
	}

	current := ""
	t.token = ""

	i := t.pos
	for i < len(t.runes) {
		current += string(t.runes[i])
		i++
		if isValidToken(current) {
			t.pos = i
			t.token = current
			t.done = t.pos >= len(t.runes)
		}
	}

	if t.token == "" {
		t.err = errors.New("piqad.Tokenizer: string contains invalid symbols and cannot be parsed well")
		return false
	}

	return true
}

// Text returns a last found token
func (t *Tokenizer) Text() string {
	return t.token
}

func isValidToken(token string) bool {
	_, ok := piqadMap[token]
	return ok
}

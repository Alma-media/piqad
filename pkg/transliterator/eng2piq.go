package transliterator

// Eng2PiqTransliterator implements a transliterator from english to klingon alphabet
type Eng2PiqTransliterator struct {
	pos   int
	runes []rune
}

// NewEng2PiqTransliterator returns a new instance of a transliterator
func NewEng2PiqTransliterator(input string) *Eng2PiqTransliterator {
	runes := []rune(input)
	return &Eng2PiqTransliterator{
		runes: runes,
		pos:   0,
	}
}

// Next increases a cursor position and returns true if it's possible to continue iterate
// Pay attention that it's not thread-safe function
func (t *Eng2PiqTransliterator) Next() bool {
	if t.pos >= len(t.runes) {
		return false
	}

	t.pos++
	return true
}

// Get returns a pIqaD letter name matched to an english symbol under a cursor position
func (t *Eng2PiqTransliterator) Get() (string, error) {
	r := t.runes[t.pos]
	return Eng2Piq(r)
}

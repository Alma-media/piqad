package piqad

import (
	"reflect"
	"testing"
)

func TestNewEng2PiqTransliterator(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want *EnglishTransliterator
	}{
		{
			"Builds a structure from an input string",
			args{"Hello"},
			&EnglishTransliterator{
				pos:   0,
				runes: []rune("Hello"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEng2PiqTransliterator(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEng2PiqTransliterator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEng2PiqTransliterator_Next(t *testing.T) {
	tests := []struct {
		name           string
		transliterator *EnglishTransliterator
		want           bool
	}{
		{
			"Returns false if the input is empty",
			&EnglishTransliterator{pos: 0, runes: []rune{}},
			false,
		},

		{
			"Returns false if cursor is out of range",
			&EnglishTransliterator{pos: 1, runes: []rune{'a'}},
			false,
		},

		{
			"Returns true if cursor is not out of range",
			&EnglishTransliterator{pos: 0, runes: []rune{'a'}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.transliterator.Next(); got != tt.want {
				t.Errorf("Eng2PiqTransliterator.Next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEng2PiqTransliterator_Get(t *testing.T) {
	tests := []struct {
		name           string
		transliterator *EnglishTransliterator
		want           string
		wantErr        bool
	}{
		{
			"Returns error if symbol cannot be matched",
			&EnglishTransliterator{pos: 0, runes: []rune("?")},
			"",
			true,
		},

		{
			"Returns a pIqaD letter name for a lower case symbol",
			&EnglishTransliterator{pos: 1, runes: []rune("axc")},
			"tlh",
			false,
		},

		{
			"Returns a pIqaD letter name for a upper case symbol",
			&EnglishTransliterator{pos: 0, runes: []rune("Xac")},
			"tlh",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.transliterator.Get()
			if (err != nil) != tt.wantErr {
				t.Errorf("Eng2PiqTransliterator.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Eng2PiqTransliterator.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEng2Piq(t *testing.T) {
	type args struct {
		eng rune
	}
	tests := [...]struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Returns a matched symbold when input symbol is downcased", args{'x'}, "tlh", false},
		{"Returns a matched symbold when input symbol is upcased", args{'X'}, "tlh", false},
		{"Returns an error when input cannot be matched", args{'?'}, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := eng2Piq(tt.args.eng)
			if (err != nil) != tt.wantErr {
				t.Errorf("Eng2Piq() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Eng2Piq() = %v, want %v", got, tt.want)
			}
		})
	}
}

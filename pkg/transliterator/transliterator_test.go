package transliterator

import (
	"testing"
)

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
			got, err := Eng2Piq(tt.args.eng)
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

func TestPiq2UTF8Code(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"Return utf8 code for a simple symbol", args{"a"}, 0xF8D0, false},
		{"Return utf8 code for a complex symbol", args{"tlh"}, 0xF8E5, false},
		{"Return an error when cannot match symbol", args{"?"}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Piq2UTF8Code(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Piq2UTF8Code() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Piq2UTF8Code() = %v, want %v", got, tt.want)
			}
		})
	}
}

package piqad

import (
	"reflect"
	"testing"
)

func TestNewPiqTokenizer(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want *Tokenizer
	}{
		{
			"Returns a new tokenizer",
			args{"ng"},
			&Tokenizer{runes: []rune("ng"), maxTokenSize: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPiqTokenizer(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPiqTokenizer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPiqTokenizer_Next(t *testing.T) {
	t.Run("Founds all tokens when whole input is valid", func(t *testing.T) {
		tokenizer := NewPiqTokenizer("ttlhnngtlht")
		want := []string{"t", "tlh", "n", "ng", "tlh", "t"}
		got := make([]string, 0)

		for tokenizer.Next() {
			token, _ := tokenizer.Get()
			got = append(got, token)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("NewPiqTokenizer() = %v, want %v", got, want)
		}
	})

	t.Run("Found valid tokens and interrups on a first invalid", func(t *testing.T) {
		tokenizer := NewPiqTokenizer("ttlhn?ng")
		want := []string{"t", "tlh", "n"}
		got := make([]string, 0)

		for tokenizer.Next() {
			token, _ := tokenizer.Get()
			got = append(got, token)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("NewPiqTokenizer() = %v, want %v", got, want)
		}
	})
}

func TestPiqTokenizer_Get(t *testing.T) {
	type fields struct {
		token string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{"Returns a last found token", fields{"tlh"}, "tlh", false},
		{"Returns an error when there is no token", fields{""}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tok := &Tokenizer{
				token: tt.fields.token,
			}
			got, err := tok.Get()
			if (err != nil) != tt.wantErr {
				t.Errorf("PiqTokenizer.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PiqTokenizer.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

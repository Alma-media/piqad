package piqad

import (
	"errors"
	"reflect"
	"testing"
)

func TestNewTokenizer(t *testing.T) {
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
			&Tokenizer{runes: []rune("ng")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTokenizer(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPiqTokenizer() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestTokenizer_Next(t *testing.T) {
// 	t.Run("Founds all tokens when whole input is valid", func(t *testing.T) {
// 		tokenizer := NewTokenizer("ttlhnngtlht")
// 		want := []string{"t", "tlh", "n", "ng", "tlh", "t"}
// 		got := make([]string, 0)

// 		for tokenizer.Scan() {
// 			token, _ := tokenizer.Text()
// 			got = append(got, token)
// 		}

// 		if !reflect.DeepEqual(got, want) {
// 			t.Errorf("NewPiqTokenizer() = %v, want %v", got, want)
// 		}
// 	})

// 	t.Run("Found valid tokens and interrups on a first invalid", func(t *testing.T) {
// 		tokenizer := NewTokenizer("ttlhn?ng")
// 		want := []string{"t", "tlh", "n"}
// 		got := make([]string, 0)

// 		for tokenizer.Scan() {
// 			token, _ := tokenizer.Text()
// 			got = append(got, token)
// 		}

// 		if !reflect.DeepEqual(got, want) {
// 			t.Errorf("NewPiqTokenizer() = %v, want %v", got, want)
// 		}
// 	})
// }

// func TestTokenizer_Text(t *testing.T) {
// 	type fields struct {
// 		token string
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		want    string
// 		wantErr bool
// 	}{
// 		{"Returns a last found token", fields{"tlh"}, "tlh", false},
// 		{"Returns an error when there is no token", fields{""}, "", true},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tok := &Tokenizer{
// 				token: tt.fields.token,
// 			}
// 			got, err := tok.Text()
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Tokenizer.Text() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if got != tt.want {
// 				t.Errorf("Tokenizer.Text() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestTokenizer_Err(t *testing.T) {
	type fields struct {
		err error
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"Returns an error", fields{errors.New("error")}, true},
		{"Does not return an error", fields{nil}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tok := &Tokenizer{
				err: tt.fields.err,
			}
			if err := tok.Err(); (err != nil) != tt.wantErr {
				t.Errorf("Tokenizer.Err() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTokenizer_Text(t *testing.T) {
	type fields struct {
		token string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Returns a found token", fields{"tlh"}, "tlh"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tok := &Tokenizer{
				token: tt.fields.token,
			}
			if got := tok.Text(); got != tt.want {
				t.Errorf("Tokenizer.Text() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenizer_Scan(t *testing.T) {
	t.Run("Founds all tokens when whole input is valid", func(t *testing.T) {
		tokenizer := NewTokenizer("ttlhnngtlht")
		want := []string{"t", "tlh", "n", "ng", "tlh", "t"}
		got := make([]string, 0)

		for tokenizer.Scan() {
			token := tokenizer.Text()
			got = append(got, token)
		}

		if err := tokenizer.Err(); err != nil {
			t.Errorf("NewPiqTokenizer.Next() is not expected to return an error, but '%v'", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Tokenizer.Next() is expected to found %v, but found %v", want, got)
		}
	})

	t.Run("Found valid tokens and interrups on a first invalid", func(t *testing.T) {
		tokenizer := NewTokenizer("ttlhn?ng")
		want := []string{"t", "tlh", "n"}
		got := make([]string, 0)

		for tokenizer.Scan() {
			token := tokenizer.Text()
			got = append(got, token)
		}

		if err := tokenizer.Err(); err == nil {
			t.Errorf("NewPiqTokenizer.Next() is expected to return an error, but has not")
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Tokenizer.Next() is expected to found %v, but found %v", want, got)
		}
	})
}

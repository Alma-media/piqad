package piqad

import (
	"testing"
)

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
		{"Return utf8 code for a complex symbol", args{"tlh"}, 0xF8E4, false},
		{"Return an error when cannot match symbol", args{"?"}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUTF8(tt.args.input)
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

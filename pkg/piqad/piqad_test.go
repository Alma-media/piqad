package piqad

import "testing"

func TestTransliterate(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Returns a transliterated string", args{"xmass"}, "tlhmaSS", false},
		{"Returns a transliterated string", args{"XMASS"}, "tlhmaSS", false},
		{"Returns a transliterated string", args{"1234567890,."}, "1234567890,.", false},
		{"Returns a transliterated string", args{"xmass, konan"}, "tlhmaSS, qonan", false},
		{"Returns an empty string", args{""}, "", false},
		{"Returns an error", args{"xmass!"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Transliterate(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("Transliterate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Transliterate() = %v, want %v", got, tt.want)
			}
		})
	}
}

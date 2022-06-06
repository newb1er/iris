package sqlx

import (
	"fmt"
	"testing"
)

func Test_snakeCase(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "OAoo",
			want:  "o_aoo",
		},
		{
			input: "OAOl",
			want:  "oaol",
		},
		{
			input: "A",
			want:  "a",
		},
		{
			input: "oAo",
			want:  "o_ao",
		},
	}
	for idx, tt := range tests {
		t.Run(fmt.Sprintf("%s[%d]", t.Name(), idx), func(t *testing.T) {
			if got := snakeCase(tt.input); got != tt.want {
				t.Errorf("snakeCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isUppercase(t *testing.T) {
	tests := []struct {
		name string
		c    rune
		want bool
	}{
		{
			name: "PC: TT",
			c:    'A',
			want: true,
		},
		{
			name: "PC: FF",
			c:    'a',
			want: false,
		},
		{
			name: "CACC: FT",
			c:    'a' - 1,
			want: false,
		},
		{
			name: "CACC: TF",
			c:    'z' + 1,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUppercase(tt.c); got != tt.want {
				t.Errorf("isUppercase() = %v, want %v", got, tt.want)
			}
		})
	}
}

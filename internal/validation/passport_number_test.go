package validation

import (
	"testing"
)

type testCase struct {
	data string
	want error
}

var tests = []testCase{
	{"1234 567890", nil},
	{"1234 5678901", ErrInvalidPassportNumber},
	{"1234 56789012", ErrInvalidPassportNumber},
	{"1234 56f890", ErrInvalidPassportNumber},
	{"123456f890", ErrInvalidPassportNumber},
	{"123456f8901", ErrInvalidPassportNumber},
	{"123 56f8901", ErrInvalidPassportNumber},
	{"123 56f890", ErrInvalidPassportNumber},
}

func TestPassportNumber(t *testing.T) {
	for _, tt := range tests {
		t.Run("PassportNumber", func(t *testing.T) {
			actual := PassportNumber(tt.data)
			if actual != tt.want {
				t.Errorf("PassportNumber(%q) = %v; want %v", tt.data, actual, tt.want)
			}
		})
	}
}

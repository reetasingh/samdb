package core

import (
	"errors"
	"fmt"
	"testing"
)

func TestEncodeOne(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected []byte
		err      error
	}{
		{
			name:     "Encode int64",
			input:    int64(123),
			expected: []byte(":123\r\n"),
			err:      nil,
		},
		{
			name:     "Encode string",
			input:    "hello",
			expected: []byte("+hello\r\n"),
			err:      nil,
		},
		{
			name:     "Encode error",
			input:    errors.New("some error"),
			expected: []byte("-some error\r\n"),
			err:      nil,
		},
		{
			name:     "Encode array of strings",
			input:    []string{"one", "two", "three"},
			expected: []byte("*3\r\n+one\r\n+two\r\n+three\r\n"),
			err:      nil,
		},
		{
			name:     "Encode empty array",
			input:    []string{},
			expected: []byte("*0\r\n"),
			err:      nil,
		},
		{
			name:     "Encode unknown type",
			input:    3.14, // Unsupported type (float64 in this case)
			expected: []byte{},
			err:      fmt.Errorf("wrong type"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := EncodeOne(tt.input)
			if string(result) != string(tt.expected) {
				t.Errorf("EncodeOne(%v) = %s; want %s", tt.input, result, tt.expected)
			}
			if err != nil && err.Error() != tt.err.Error() {
				t.Errorf("EncodeOne(%v) error = %v; want %v", tt.input, err, tt.err)
			} else if err == nil && tt.err != nil {
				t.Errorf("EncodeOne(%v) error = nil; want %v", tt.input, tt.err)
			}
		})
	}
}

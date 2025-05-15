package helpers

import (
	"testing"
)

func TestEncodeBase64(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Simple string",
			input:    "hello",
			expected: "aGVsbG8=",
		},
		{
			name:     "Complex string",
			input:    "Hello, World! 123",
			expected: "SGVsbG8sIFdvcmxkISAxMjM=",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EncodeBase64(tt.input)
			if result != tt.expected {
				t.Errorf("EncodeBase64(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

package main

import "testing"

func TestContainsControlRune(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"valid string", "hello world", false},
		{"string with newline", "hello\nworld", true},
		{"string with slash", "hello/world", false},
		{"string with null terminating symbol", "hello\x00world", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsControlRune(tt.input)

			if got != tt.want {
				t.Errorf("expected %v, got %v", tt.want, got)
			}
		})
	}
}

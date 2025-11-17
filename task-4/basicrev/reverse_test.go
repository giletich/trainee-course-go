package basicreverse

import "testing"

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BasicReverse("bobob bskdf bovov")
	}
}

func TestIsPalindromeThird(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "1",
			input:    "яяя",
			expected: "яяя",
		},
		{
			name:     "2",
			input:    "abo ba",
			expected: "oba ab",
		},
		{
			name:     "3",
			input:    "asdf asdf r r",
			expected: "fdsa fdsa r r",
		},
		{
			name:     "4",
			input:    "hello мир",
			expected: "olleh рим",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BasicReverse(tt.input)
			if result != tt.expected {
				t.Errorf("pal3(%q) = %v; expected %v",
					tt.input, result, tt.expected)
			}
		})
	}
}

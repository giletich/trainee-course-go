package palindrome

import "testing"

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PalindromeFirst("bobobob")
	}
}

func BenchmarkSecond(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PalindromeSecond("bobobob")
	}
}

func BenchmarkThird(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PalindromeThird("bobobob", 0, 6)
	}
}

func TestIsPalindromeFirst(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected bool
    }{
        {
            name:     "1",
            input:    "я",
            expected: true,
        },
        {
            name:     "2",
            input:    "яяя яяя",
            expected: true,
        },
        {
            name:     "3",
            input:    "мяу",
            expected: false,
        },
        {
            name:     "4",
            input:    "hello мир",
            expected: false,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := PalindromeFirst(tt.input)
            if result != tt.expected {
                t.Errorf("pal1(%q) = %v; expected %v", 
                    tt.input, result, tt.expected)
            }
        })
    }
}

func TestIsPalindromeSecond(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected bool
    }{
        {
            name:     "1",
            input:    "я",
            expected: true,
        },
        {
            name:     "2",
            input:    "яяя яяя",
            expected: true,
        },
        {
            name:     "3",
            input:    "мяу",
            expected: false,
        },
        {
            name:     "4",
            input:    "hello мир",
            expected: false,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := PalindromeFirst(tt.input)
            if result != tt.expected {
                t.Errorf("pal2(%q) = %v; expected %v", 
                    tt.input, result, tt.expected)
            }
        })
    }
}

func TestIsPalindromeThird(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected bool
    }{
        {
            name:     "1",
            input:    "я",
            expected: true,
        },
        {
            name:     "2",
            input:    "яяя яяя",
            expected: true,
        },
        {
            name:     "3",
            input:    "мяу",
            expected: false,
        },
        {
            name:     "4",
            input:    "hello мир",
            expected: false,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := PalindromeThird(tt.input, 0, len(tt.input)-1)
            if result != tt.expected {
                t.Errorf("pal3(%q) = %v; expected %v", 
                    tt.input, result, tt.expected)
            }
        })
    }
}

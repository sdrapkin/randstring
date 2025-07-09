package randstring

import (
	cryptoRand "crypto/rand"
	"fmt"
	"runtime"
	"testing"
	"unsafe"
)

func TestTextFunctions(t *testing.T) {
	tests := []struct {
		name     string
		fn       func(int) string
		alphabet string
	}{
		{"Text16", Text16, base16},
		{"Text32", Text32, base32},
		{"Text32c", Text32c, base32c},
		{"Text64", Text64, base64},
		{"Text64URL", Text64URL, base64url},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test length and charset
			s := tt.fn(64)
			if len(s) != 64 {
				t.Errorf("expected length 64, got %d", len(s))
			}
			for _, c := range s {
				if !containsRune(tt.alphabet, c) {
					t.Errorf("invalid character %q in %s", c, tt.name)
				}
			}

			// Test zero length
			if got := tt.fn(0); got != "" {
				t.Errorf("%s(0): expected empty string, got %q", tt.name, got)
			}

			// Test negative length panics
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("%s(-1): expected panic for negative length, but did not panic", tt.name)
				}
			}()
			tt.fn(-1)
		})
	}
}

func TestUniqueness(t *testing.T) {
	tests := []struct {
		name string
		fn   func(int) string
	}{
		{"Text16", Text16},
		{"Text32", Text32},
		{"Text32c", Text32c},
		{"Text64", Text64},
		{"Text64URL", Text64URL},
	}
	const ITERATIONS = 10000
	const STRING_LENGTH = 32
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seen := make(map[string]struct{}, ITERATIONS)
			for range ITERATIONS {
				s := tt.fn(STRING_LENGTH)
				if _, ok := seen[s]; ok {
					t.Fatalf("%s: duplicate string generated: %q", tt.name, s)
				}
				seen[s] = struct{}{}
			}
		})
	}
}

// Helper function to check if a rune is in a string
func containsRune(s string, r rune) bool {
	for _, c := range s {
		if c == r {
			return true
		}
	}
	return false
} // containsRune

func TestAllLengths(t *testing.T) {
	tests := []struct {
		name string
		fn   func(int) string
	}{
		{"Text16", Text16},
		{"Text32", Text32},
		{"Text32c", Text32c},
		{"Text64", Text64},
		{"Text64URL", Text64URL},
	}
	const maxLen = 1024
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for length := 0; length <= maxLen; length++ {
				s := tt.fn(length)
				if len(s) != length {
					t.Errorf("%s(%d): expected length %d, got %d", tt.name, length, length, len(s))
				}
			}
		})
	}
} // TestAllLengths

func FuzzTextFunctions(f *testing.F) {
	funcs := []struct {
		name string
		fn   func(int) string
	}{
		{"Text16", Text16},
		{"Text32", Text32},
		{"Text32c", Text32c},
		{"Text64", Text64},
		{"Text64URL", Text64URL},
	}
	// Seed with interesting values
	for _, length := range []int{-100, -1, 0, 1, 2, 16, 127, 128, 129, 255, 256, 512, 1024, 4096} {
		f.Add(length)
	}
	f.Fuzz(func(t *testing.T, length int) {
		for _, fn := range funcs {
			if length < 0 {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("%s: expected panic for negative length %d", fn.name, length)
					}
				}()
				fn.fn(length)
			} else {
				s := fn.fn(length)
				if len(s) != length {
					t.Errorf("%s(%d): expected length %d, got %d", fn.name, length, length, len(s))
				}
			}
		}
	})
} // FuzzTextFunctions

func TestUnsafeStringGC(t *testing.T) {
	const ITERATIONS = 100_000
	strings := make([]string, ITERATIONS)
	for i := range ITERATIONS {
		byteSlice := fmt.Appendf(nil, "%d", i)
		strings[i] = unsafe.String(&byteSlice[0], len(byteSlice))
	}
	runtime.GC()
	for i := range ITERATIONS {
		expected := fmt.Sprintf("%d", i)
		if strings[i] != expected {
			t.Fatalf("GC test failed at index %d: got %q, want %q", i, strings[i], expected)
		}
	}
} // TestUnsafeStringGC

// ===============================================================================
// Benchmarks for the Text functions. go test -bench=".*" -benchmem -benchtime=5s
// ===============================================================================

// Benchmark for crypto/rand.Text (Go 1.24+)
func BenchmarkCryptoRandText_StdLib(b *testing.B) {
	for b.Loop() {
		_ = cryptoRand.Text()
	}
}
func BenchmarkText16_32chars(b *testing.B) {
	for b.Loop() {
		_ = Text16(32)
	}
}
func BenchmarkText16_1024chars(b *testing.B) {
	for b.Loop() {
		_ = Text16(1024)
	}
}

func BenchmarkText32_32chars(b *testing.B) {
	for b.Loop() {
		_ = Text32(512)
	}
}

func BenchmarkText32_1024chars(b *testing.B) {
	for b.Loop() {
		_ = Text32(1024)
	}
}

func BenchmarkText64_32chars(b *testing.B) {
	for b.Loop() {
		_ = Text64(32)
	}
}

func BenchmarkText64_1024chars(b *testing.B) {
	for b.Loop() {
		_ = Text64(1024)
	}
}

func BenchmarkText64URL_32chars(b *testing.B) {
	for b.Loop() {
		_ = Text64URL(32)
	}
}

func BenchmarkText64URL_1024chars(b *testing.B) {
	for b.Loop() {
		_ = Text64URL(1024)
	}
}

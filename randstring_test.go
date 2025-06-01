package randstring

import (
	cryptoRand "crypto/rand"
	"testing"
)

func TestText16_LengthAndCharset(t *testing.T) {
	const alphabet = "0123456789ABCDEF"
	s := Text16(64)
	if len(s) != 64 {
		t.Errorf("Text16: expected length 64, got %d", len(s))
	}
	for _, c := range s {
		if !containsRune(alphabet, c) {
			t.Errorf("Text16: invalid character %q", c)
		}
	}
} // TestText16_LengthAndCharset

func TestText32_LengthAndCharset(t *testing.T) {
	const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"
	s := Text32(64)
	if len(s) != 64 {
		t.Errorf("Text32: expected length 64, got %d", len(s))
	}
	for _, c := range s {
		if !containsRune(alphabet, c) {
			t.Errorf("Text32: invalid character %q", c)
		}
	}
} // TestText32_LengthAndCharset

func TestText64_LengthAndCharset(t *testing.T) {
	const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	s := Text64(64)
	if len(s) != 64 {
		t.Errorf("Text64: expected length 64, got %d", len(s))
	}
	for _, c := range s {
		if !containsRune(alphabet, c) {
			t.Errorf("Text64: invalid character %q", c)
		}
	}
} // TestText64_LengthAndCharset

func TestText64URL_LengthAndCharset(t *testing.T) {
	const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
	s := Text64URL(64)
	if len(s) != 64 {
		t.Errorf("Text64URL: expected length 64, got %d", len(s))
	}
	for _, c := range s {
		if !containsRune(alphabet, c) {
			t.Errorf("Text64URL: invalid character %q", c)
		}
	}
} // TestText64URL_LengthAndCharset

func TestUniqueness(t *testing.T) {
	const ITERATIONS = 1000
	const STRING_LENGTH = 32
	types := []struct {
		name string
		fn   func(int) string
	}{
		{"Text16", Text16},
		{"Text32", Text32},
		{"Text64", Text64},
		{"Text64URL", Text64URL},
	}
	for _, typ := range types {
		seen := make(map[string]struct{}, ITERATIONS)
		for range ITERATIONS {
			s := typ.fn(STRING_LENGTH)
			if _, exists := seen[s]; exists {
				t.Errorf("%s: duplicate string generated: %q", typ.name, s)
			}
			seen[s] = struct{}{}
		}
	}
} // TestUniqueness

func TestZeroLength(t *testing.T) {
	if s := Text16(0); s != "" {
		t.Errorf("Text16(0): expected empty string, got %q", s)
	}
	if s := Text32(0); s != "" {
		t.Errorf("Text32(0): expected empty string, got %q", s)
	}
	if s := Text64(0); s != "" {
		t.Errorf("Text64(0): expected empty string, got %q", s)
	}
	if s := Text64URL(0); s != "" {
		t.Errorf("Text64URL(0): expected empty string, got %q", s)
	}
} // TestZeroLength

func TestNegativeLengthPanics(t *testing.T) {
	tests := []struct {
		name string
		fn   func(int) string
	}{
		{"Text16", Text16},
		{"Text32", Text32},
		{"Text64", Text64},
		{"Text64URL", Text64URL},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("%s: expected panic for negative length, but did not panic", tt.name)
				}
			}()
			tt.fn(-1)
		})
	}
} // TestNegativeLengthPanics

// Helper function to check if a rune is in a string
func containsRune(s string, r rune) bool {
	for _, c := range s {
		if c == r {
			return true
		}
	}
	return false
} // containsRune

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
		_ = Text32(32)
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

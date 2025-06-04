// randstring provides functions to generate random strings using various encoding alphabets.
// It extends the standard library's crypto/rand .Text() function. (https://pkg.go.dev/crypto/rand#Text)
package randstring

import (
	cryptoRand "crypto/rand"
	"unsafe"
)

const (
	// Standard Base16/Base32/Base64 encoding alphabets from RFC 4648.
	base16    = "0123456789ABCDEF"
	base32    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"
	base64    = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	base64url = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"

	// base16_256 is the base16 repeated enough times to cover all byte values (0-255).
	base16_256 = base16 + base16 + base16 + base16 +
		base16 + base16 + base16 + base16 +
		base16 + base16 + base16 + base16 +
		base16 + base16 + base16 + base16

	// base32_256 is the base32 repeated enough times to cover all byte values (0-255).
	base32_256 = base32 + base32 + base32 + base32 +
		base32 + base32 + base32 + base32

	// base64_256 is the base64 repeated enough times to cover all byte values (0-255).
	base64_256 = base64 + base64 + base64 + base64

	// base64url_256 is the base64url repeated enough times to cover all byte values (0-255).
	base64url_256 = base64url + base64url + base64url + base64url
)

func textAlphabet256(alphabet256 string, length int) string {
	if length == 0 {
		return ""
	}
	buffer := make([]byte, length) // dynamic-length make will cause buffer to be heap-allocated
	cryptoRand.Read(buffer)        // guaranteed not to fail since Go 1.24

	for i := range buffer {
		buffer[i] = alphabet256[buffer[i]]
	}
	return unsafe.String(&buffer[0], length)
} // textAlphabet256

// Text16 generates a random string of the specified length using the Base16 alphabet.
func Text16(length int) string {
	return textAlphabet256(base16_256, length)
}

// Text32 generates a random string of the specified length using the Base32 alphabet.
func Text32(length int) string {
	return textAlphabet256(base32_256, length)
}

// Text64 generates a random string of the specified length using the Base64 alphabet.
func Text64(length int) string {
	return textAlphabet256(base64_256, length)
}

// Text64URL generates a random string of the specified length using the Base64 URL-safe alphabet.
func Text64URL(length int) string {
	return textAlphabet256(base64url_256, length)
}

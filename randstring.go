// randstring provides functions to generate random strings using various encoding alphabets.
// It extends the standard library's crypto/rand .Text() function. (https://pkg.go.dev/crypto/rand#Text)
package randstring

import (
	cryptoRand "crypto/rand"
)

// Standard Base16/Base32/Base64 encoding alphabets from RFC 4648.
const base16alphabet = "0123456789ABCDEF"
const base32alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"
const base64alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
const base64url_alph = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"

// base16alphabet256 is the base16alphabet repeated enough times to cover all byte values (0-255).
const base16alphabet256 = base16alphabet + base16alphabet + base16alphabet + base16alphabet +
	base16alphabet + base16alphabet + base16alphabet + base16alphabet +
	base16alphabet + base16alphabet + base16alphabet + base16alphabet +
	base16alphabet + base16alphabet + base16alphabet + base16alphabet

// base32alphabet256 is the base32alphabet repeated enough times to cover all byte values (0-255).
const base32alphabet256 = base32alphabet + base32alphabet + base32alphabet + base32alphabet +
	base32alphabet + base32alphabet + base32alphabet + base32alphabet

// base64alphabet256 is the base64alphabet repeated enough times to cover all byte values (0-255).
const base64alphabet256 = base64alphabet + base64alphabet + base64alphabet + base64alphabet

// base64url_alph256 is the base64url_alph repeated enough times to cover all byte values (0-255).
const base64url_alph256 = base64url_alph + base64url_alph + base64url_alph + base64url_alph

func textAlphabet256(alphabet256 string, length int) string {
	const MAX_STACK_ALLOC = 128
	var buffer []byte
	if length <= MAX_STACK_ALLOC {
		var stackBuffer [MAX_STACK_ALLOC]byte
		buffer = stackBuffer[:length]
	} else {
		buffer = make([]byte, length)
	}
	cryptoRand.Read(buffer) // guaranteed not to fail since Go 1.24
	for i := range buffer {
		buffer[i] = alphabet256[buffer[i]]
	}
	return string(buffer)
} // textAlphabet256

// Text16 generates a random string of the specified length using the Base16 alphabet.
func Text16(length int) string {
	return textAlphabet256(base16alphabet256, length)
}

// Text32 generates a random string of the specified length using the Base32 alphabet.
func Text32(length int) string {
	return textAlphabet256(base32alphabet256, length)
}

// Text64 generates a random string of the specified length using the Base64 alphabet.
func Text64(length int) string {
	return textAlphabet256(base64alphabet256, length)
}

// Text64URL generates a random string of the specified length using the Base64 URL-safe alphabet.
func Text64URL(length int) string {
	return textAlphabet256(base64url_alph256, length)
}

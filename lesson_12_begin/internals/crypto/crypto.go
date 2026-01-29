package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

// NonceSize is the size of the GCM nonce
const NonceSize = 12

// DeriveKey derives a 32-byte AES-256 key from the shared secret
func DeriveKey(secret string) []byte {
	// TODO: Create hash by calling sha256.Sum256(), pass []byte(secret) as argument
	// return hash[:]

}

// Encrypt encrypts plaintext using AES-GCM and returns base64-encoded result
func Encrypt(plaintext []byte, secret string) (string, error) {
	key := DeriveKey(secret)

	// Create AES cipher
	// TODO: create block by calling aes.NewCipher(), pass key as argument
	if err != nil {
		return "", fmt.Errorf("creating cipher: %w", err)
	}

	// Create GCM mode
	// TODO: create gcm by calling cipher.NewGCM, pass block as argument
	if err != nil {
		return "", fmt.Errorf("creating GCM: %w", err)
	}

	// Generate random nonce
	nonce := make([]byte, NonceSize)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("generating nonce: %w", err)
	}

	// Encrypt and append tag
	// TODO: create ciphertext by calling gcm.Seal()

	// Prepend nonce to ciphertext
	// TODO: use append to present nonce to ciphertext, assign to result

	// Base64 encode for transmission
	return base64.StdEncoding.EncodeToString(result), nil
}

// Decrypt decrypts base64-encoded ciphertext using AES-GCM
func Decrypt(encoded string, secret string) ([]byte, error) {
	key := DeriveKey(secret)

	// Base64 decode
	data, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, fmt.Errorf("base64 decode: %w", err)
	}

	// Check minimum length (nonce + at least some ciphertext)
	if len(data) < NonceSize+1 {
		return nil, fmt.Errorf("ciphertext too short")
	}

	// Extract nonce and ciphertext
	// TODO: assign nonce to data[:NonceSize]
	// TODO: assign ciphertext to data[NonceSize:]

	// Create AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("creating cipher: %w", err)
	}

	// Create GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("creating GCM: %w", err)
	}

	// Decrypt and verify tag
	// TODO: call gcm.Open(), assign return value to plaintext
	
	if err != nil {
		return nil, fmt.Errorf("decryption failed: %w", err)
	}

	return plaintext, nil
}

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
	hash := sha256.Sum256([]byte(secret))
	return hash[:]
}

// Encrypt encrypts plaintext using AES-GCM and returns base64-encoded result
func Encrypt(plaintext []byte, secret string) (string, error) {
	key := DeriveKey(secret)

	// Create AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("creating cipher: %w", err)
	}

	// Create GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("creating GCM: %w", err)
	}

	// Generate random nonce
	nonce := make([]byte, NonceSize)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("generating nonce: %w", err)
	}

	// Encrypt and append tag
	ciphertext := gcm.Seal(nil, nonce, plaintext, nil)

	// Prepend nonce to ciphertext
	result := append(nonce, ciphertext...)

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
	nonce := data[:NonceSize]
	ciphertext := data[NonceSize:]

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
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("decryption failed: %w", err)
	}

	return plaintext, nil
}

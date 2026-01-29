package server

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

const (
	// TimestampTolerance is how far off the timestamp can be (in seconds)
	TimestampTolerance = 300 // 5 minutes
)

// VerifyRequest checks HMAC signature and timestamp validity
func VerifyRequest(r *http.Request, secret string) error {
	// Extract headers
	timestamp := r.Header.Get("X-Auth-Timestamp")
	signature := r.Header.Get("X-Auth-Signature")

	if timestamp == "" || signature == "" {
		return fmt.Errorf("missing authentication headers")
	}

	// Verify timestamp is within tolerance
	if err := verifyTimestamp(timestamp); err != nil {
		return err
	}

	// Read the body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("reading body: %w", err)
	}

	// Restore the body for downstream handlers
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	// Recompute the signature
	message := timestamp + string(body)
	expectedSignature := serverComputeHMAC(message, secret)

	// Constant-time comparison to prevent timing attacks
	if !hmac.Equal([]byte(signature), []byte(expectedSignature)) {
		return fmt.Errorf("invalid signature")
	}

	return nil
}

// verifyTimestamp checks if timestamp is within acceptable range
func verifyTimestamp(timestampStr string) error {
	timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid timestamp format")
	}

	now := time.Now().Unix()
	diff := now - timestamp

	// Check if timestamp is too old or too far in the future
	if diff < -TimestampTolerance || diff > TimestampTolerance {
		return fmt.Errorf("timestamp outside acceptable range")
	}

	return nil
}

// serverComputeHMAC calculates HMAC-SHA256 (same as agent)
func serverComputeHMAC(message, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(message))
	return hex.EncodeToString(mac.Sum(nil))
}

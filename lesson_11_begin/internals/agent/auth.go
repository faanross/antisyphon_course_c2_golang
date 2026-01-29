package agent

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"strconv"
	"time"
)

// SignRequest adds HMAC authentication headers to an HTTP request
func SignRequest(req *http.Request, body []byte, secret string) {
	// Get current timestamp
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	// Create the message to sign: timestamp + body
	// TODO: Create message equal to timestamp (for randomness) + body cast to string

	// Compute HMAC-SHA256
	// TODO: Create signature by calling computeHMAC(), pass message and secret

	// Add headers
	req.Header.Set("X-Auth-Timestamp", timestamp)
	req.Header.Set("X-Auth-Signature", signature)
}

// computeHMAC calculates HMAC-SHA256 and returns hex-encoded result
func computeHMAC(message, secret string) string {
	// TODO: Calculate mac by calling hmac.New(), pass sha256.New and []byte(secret) as arguments
	// TODO: Call mac.Write(), pass []byte(message) as argument
	return hex.EncodeToString(mac.Sum(nil))
}

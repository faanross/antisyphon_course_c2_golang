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
	message := timestamp + string(body)

	// Compute HMAC-SHA256
	signature := computeHMAC(message, secret)

	// Add headers
	req.Header.Set("X-Auth-Timestamp", timestamp)
	req.Header.Set("X-Auth-Signature", signature)
}

// computeHMAC calculates HMAC-SHA256 and returns hex-encoded result
func computeHMAC(message, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(message))
	return hex.EncodeToString(mac.Sum(nil))
}

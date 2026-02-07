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

// TimestampTolerance is how far off the timestamp can be (in seconds)
// TODO: define constant TimestampTolerance as 5 mins

// VerifyRequest checks HMAC signature and timestamp validity
func VerifyRequest(r *http.Request, secret string) error {
	// Extract headers

	// TODO: extract timestamp from X-Auth-Timestamp header
	// TODO extract signature from X-Auth-Signature header

	// TODO if either timestamp or signature is blank - return error

	// Verify timestamp is within tolerance
	// TODO: call verifyTimestamp()

	// Read the body
	// TODO read r.Body

	// Recompute the signature
	// TODO set message equal to timestamp + string(body)
	// TODO: call computeHMAC(), returns expectedSignature

	// Constant-time comparison to prevent timing attacks
	// TODO: Perform constant-time comparison between signature and expectedSignature

	return nil
}

// verifyTimestamp checks if timestamp is within acceptable range
func verifyTimestamp(timestampStr string) error {

	// TODO: calculate timestamp using strconv.ParseInt()
	if err != nil {
		return fmt.Errorf("invalid timestamp format")
	}

	// TODO: set now equal to current Unix time
	// TODO: Calculate different as now - timestamp

	// Check if timestamp is too old or too far in the future
	// Use TimestampTolerance to ensure diff within acceptable range

	return nil
}

// serverComputeHMAC calculates HMAC-SHA256 (same as agent)
func serverComputeHMAC(message, secret string) string {

	// TODO: calculate mac using hmac.New()
	// TODO: call max.Write(), pass message cast to byte slice
	// TODO: return mac.Sum(nil) encoded to hex

}

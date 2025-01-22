package utils

import (
	"crypto/rand"
	"sync"
	"time"
)

var (
	otpStore = make(map[string]otpEntry)
	mu       sync.Mutex
)

type otpEntry struct {
	OTP       string
	Generated time.Time
}

// GenerateOTP generates a 6-digit OTP and stores it in memory with a timestamp
func GenerateOTP(email string) string {
	mu.Lock()
	defer mu.Unlock()

	otp := generateRandomOTP()
	otpStore[email] = otpEntry{
		OTP:       otp,
		Generated: time.Now(),
	}
	return otp
}

// VerifyOTP verifies if the provided OTP is correct and not expired for the given email
func VerifyOTP(email string, otp string) bool {
	mu.Lock()
	defer mu.Unlock()

	entry, exists := otpStore[email]
	if !exists {
		return false
	}

	if time.Since(entry.Generated) > 2*time.Minute {
		delete(otpStore, email)
		return false
	}

	if entry.OTP == otp {
		delete(otpStore, email)
		return true
	}
	return false
}

// generateRandomOTP generates a 6-digit random OTP using crypto/rand
func generateRandomOTP() string {
	const otpLength = 6
	const charset = "0123456789"

	otp := make([]byte, otpLength)
	if _, err := rand.Read(otp); err != nil {
		panic(err)
	}

	for i := 0; i < otpLength; i++ {
		otp[i] = charset[otp[i]%byte(len(charset))]
	}

	return string(otp)
}

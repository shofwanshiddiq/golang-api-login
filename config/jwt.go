package config

import (
	"log"
	"os"
	"time"
)

func GetJWTSecret() []byte {
	secret := os.Getenv("JWT_SECRET_KEY")
	if secret == "" {
		log.Fatal("JWT_SECRET_KEY is not set in environment variables")
	}
	return []byte(secret) // ✅ Returns full secret
}

func GetJWTExpirationTime() time.Duration {
	durationStr := os.Getenv("JWT_EXPIRATION_IN")
	if durationStr == "" {
		log.Println("JWT_EXPIRATION_IN is not set, defaulting to 24 hours")
		return time.Hour * 24
	}

	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		log.Println("Invalid JWT_EXPIRATION_IN format, defaulting to 24 hours:", err)
		return time.Hour * 24
	}

	return duration
}

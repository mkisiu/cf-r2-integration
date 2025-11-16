package main

import (
	"log"
	"os"
)

// Config holds all data necessary to connect to R2
// For now juest from env vars
type Config struct {
	URL string
	AccountID string
	AccessKey string
	SecretKey string
	Bucket string
}

// getentRequired - helper, wchich requires thejjkk env var to be set
// If env var is not present, it ends with a msg
func getenvRequired(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok || value == "" {
		log.Fatalf("missing required env variable %s", key)
	}

	return value
}

func loadConfig() Config {
	return Config{
		URL: getenvRequired("R2_URL"),
		AccountID: getenvRequired("R2_ACCOUNT_ID"),
		AccessKey: getenvRequired("R2_ACCESS_KEY_ID"),
		SecretKey: getenvRequired("R2_SECRET_ACCESS_KEY"),
		Bucket: getenvRequired("R2_BUCKET_NAME"),
	}
}

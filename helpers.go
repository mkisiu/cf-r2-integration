package main

import (
	"fmt"
)

// maskSecret - helper, masks sensitive data
func maskSecret(s string) string {
	if len(s) <= 4 {
		return "****"
	}
	return "****" + s[len(s)-4:]
}

// showConfig - helper, shows all the configuration params
func showConfig() {
	// show loaded env vars values
	fmt.Println("Cloudflare R2 configuration loaded:")
	fmt.Printf("  URL:       %s\n", cfg.URL)
	fmt.Printf("  AccountID: %s\n", cfg.AccountID)
	fmt.Printf("  AccessKey: %s\n", cfg.AccessKey)
	fmt.Printf("  Bucket:    %s\n", cfg.Bucket)
	fmt.Printf("  SecretKey: %s...\n", maskSecret(cfg.SecretKey))
}

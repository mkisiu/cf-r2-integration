package main

import (
	"context"
	"log"
	"time"

	"github.com/mkisiu/cloudflare-r2-go"
)

func setup(ctx context.Context) {
	// load local config
	{
		cfg = loadConfig()
	}

	// configure the r2 config base on the local config
	{
		r2Config = r2.Config{
			Endpoint:        cfg.URL,
			AccessKeyID:     cfg.AccessKey,
			SecretAccessKey: cfg.SecretKey,
			BucketName:      cfg.Bucket,
		}
	}

	// create a new r2 client with the r2Config
	{
		// timeout for creation the R2 client
		clientCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		var err error
		r2Client, err = r2.NewClient(clientCtx, r2Config)
		if err != nil {
			log.Fatalf("unable to create R2 client: %v", err)
		}
	}
}

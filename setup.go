package main

import (
	"context"
	"log"
	"github.com/mkisiu/cf-r2-integration/r2"
)

func setup(ctx context.Context) {
	// load local config
	{
		cfg = loadConfig()
	}

	// configure the r2 config base on the local config
	{
		r2Config.Endpoint = cfg.URL
		r2Config.AccessKeyID = cfg.AccessKey
		r2Config.SecretAccessKey = cfg.SecretKey
		r2Config.BucketName = cfg.Bucket
	}

	// create a new r2 client with the r2Config
	{
		var err error
		r2Client, err = r2.NewClient(ctx, r2Config)
		if err != nil {
			log.Fatalf("unable to create R2 client: %v", err)
		}
	}
}


package main

import (
	"context"
	"log"

	"github.com/mkisiu/cf-r2-integration/r2"
)

var (
	cfg Config
	r2Config r2.Config
	r2Client *r2.Client
)

func main() {
	// create a context
	ctx := context.Background()

	// read env vars
	setup(ctx)

	// show config params
	showConfig()

	// list objects in the bucket
	keys, err := r2Client.ListObjects(ctx, "")
	if err != nil {
		log.Fatalf("R2: ListObjects error: %v", err)
	}

	// show received list of objects (keys)
	log.Println("R2: objects in the bucket:")
	for _, k := range keys {
		log.Println(" -", k)
	}
}

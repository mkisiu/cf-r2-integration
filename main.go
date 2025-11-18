package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/mkisiu/cloudflare-r2-go"
)

var (
	cfg      Config
	r2Config r2.Config
	r2Client *r2.Client
)

func main() {
	// create a context
	ctx := context.Background()

	// read env vars
	setup(ctx)

	// show config params
	// showConfig()

	// list objects in the bucket
	listCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	keys, err := r2Client.ListObjects(listCtx, "")
	if err != nil {
		log.Fatalf("R2: ListObjects error: %v", err)
	}

	// get the head of the sample object from the bucket
	getCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	obj, err := r2Client.HeadObject(getCtx, keys[2])
	if err != nil {
		if errors.Is(err, r2.ErrObjectNotFound) {
			log.Println("Object not found")
			return
		} else if errors.Is(err, context.DeadlineExceeded) {
			log.Println("Deadline exceeded")
			return
		} else {
			log.Fatalf("R2: HeadObject error: %v", err)
		}
	}
	fmt.Printf("content type: %v\n", obj.ContentType)
	fmt.Printf("content lenght: %d\n", obj.ContentLength)
	fmt.Printf("last modified: %v\n", obj.LastModified)
	fmt.Printf("ETag: %v\n", obj.ETag)
	fmt.Println("metadata:")
	for k, v := range obj.Metadata {
		fmt.Printf("\t%s: %s\n", k, v)
	}
}

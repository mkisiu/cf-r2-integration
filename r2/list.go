package r2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// ListObjects - returns the list of keys (files?) in a bucket
// Prefix can be set to "" to return all the keys
func (c *Client) ListObjects(ctx context.Context, prefix string) ([]string, error) {
	// check if we received the properly configured client to fail early
	if c == nil || c.s3 == nil {
		return nil, fmt.Errorf("r2: client not initialized")
	}

	// create an object to store the results
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(c.bucket),
	}
	if prefix != "" {
		input.Prefix = aws.String(prefix)
	}

	// get the objects and put it into "result" var 
	result, err := c.s3.ListObjectsV2(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("r2: list objects: %w", err)
	}

	// create and return the list of keys getheres from the result object
	var keys []string
	for _, obj := range result.Contents {
		if obj.Key != nil {
			keys = append(keys, *obj.Key)
		}
	}

	return keys, nil
}

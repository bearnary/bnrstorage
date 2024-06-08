package bnrstorage

import (
	"context"
)

type Client interface {
	SelectBucket(ctx context.Context, bucketName string) (Bucket, error)
}

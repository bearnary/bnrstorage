package bnrstorage

import (
	"context"
	"fmt"
)

func (c *defaultClient) SelectBucket(ctx context.Context, bucketName string) (Bucket, error) {

	b := &defaultBucket{
		bucketName: bucketName,
		region:     c.region,
		mc:         c.mc,
		ctx:        ctx,
	}

	err := c.mc.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: c.region})
	if err != nil {

		exists, errBucketExists := c.mc.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			return b, nil
		} else {
			return nil, fmt.Errorf("bucket check exist fail: %v", err.Error())
		}
	}
	return b, nil
}

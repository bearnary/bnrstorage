package bnrstorage

import (
	"fmt"

	"github.com/minio/minio-go/v7"
)

func (c *defaultClient) SelectBucket(bucketName string) (Bucket, error) {

	b := &defaultBucket{
		bucketName: bucketName,
		region:     c.region,
		mc:         c.mc,
		ctx:        c.ctx,
	}

	err := c.mc.MakeBucket(c.ctx, bucketName, minio.MakeBucketOptions{Region: c.region})
	if err != nil {

		exists, errBucketExists := c.mc.BucketExists(c.ctx, bucketName)
		if errBucketExists == nil && exists {
			return b, nil
		} else {
			return nil, fmt.Errorf("bucket check exist fail: %v", err.Error())
		}
	}
	return b, nil
}

package bnrstorage

import (
	"context"
)

var mainBucketFactory *defaultMainBucketFactory

type defaultMainBucketFactory struct {
	bucket Bucket
	ctx    context.Context
}

func RegisterMainStorageBucket(ctx context.Context, client Client, storageBucketName string) error {

	b, vErr := client.SelectBucket(ctx, storageBucketName)
	if vErr != nil {
		return vErr
	}

	mainBucketFactory = &defaultMainBucketFactory{
		bucket: b,
		ctx:    ctx,
	}

	return nil
}

func MainBucketFactory() *defaultMainBucketFactory {
	if mainBucketFactory == nil {
		panic("main bucket factory not registered")
	}
	return mainBucketFactory
}

func (f *defaultMainBucketFactory) GetObjectUrl(fileName string) (string, error) {
	url, vErr := f.bucket.GetObjectUrl(f.ctx, fileName)
	return url, vErr
}

func (f *defaultMainBucketFactory) GetObjectDownloadUrl(fileName string) (string, error) {
	url, vErr := f.bucket.GetObjectDownloadUrl(f.ctx, fileName)
	return url, vErr
}

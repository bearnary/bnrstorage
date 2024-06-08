package bnrstorage

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
)

func (b *defaultBucket) DeleteObject(ctx context.Context, fileName string) error {
	opts := minio.RemoveObjectOptions{}
	err := b.mc.RemoveObject(ctx, b.bucketName, fileName, opts)
	if err != nil {
		vErr := fmt.Errorf("delete object fail: %v", err.Error())
		return vErr
	}
	return nil
}

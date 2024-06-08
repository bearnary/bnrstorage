package bnrstorage

import (
	"context"
	"fmt"
)

func (b *defaultBucket) DeleteObject(ctx context.Context, fileName string) error {
	opts := minio.RemoveObjectOptions{}
	vErr := b.mc.RemoveObject(ctx, b.bucketName, fileName, opts)
	if vErr != nil {
		vErr := fmt.Errorf("delete object fail: %v", err.Error())
		return vErr
	}
	return nil
}

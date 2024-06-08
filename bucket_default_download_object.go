package bnrstorage

import (
	"context"
	"fmt"
	"net/url"
	"time"
)

func (b *defaultBucket) GetObjectDownloadUrl(ctx context.Context, fileName string) (string, error) {

	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", "attachment; filename=\""+fileName+"\"")

	url, err := b.mc.PresignedGetObject(ctx, b.bucketName, fileName, time.Duration(10)*time.Hour, reqParams)
	if err != nil {
		return "", fmt.Errorf("get object url fail: %v", err.Error())
	}

	return url.String(), nil
}

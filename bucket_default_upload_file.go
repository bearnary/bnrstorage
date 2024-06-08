package bnrstorage

import (
	"context"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

func (b *defaultBucket) UploadFile(ctx context.Context, fileName, filePath, storagePath, contentType string) (*UploadInfo, error) {

	//Example
	//objectName := "golden-oldies.zip"
	//filePath := "/tmp/golden-oldies.zip"
	//contentType := "application/zip"

	// Upload the zip file with FPutObject
	n, err := b.mc.FPutObject(ctx, b.bucketName, storagePath+fileName, filePath, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return nil, fmt.Errorf("storage uplaod file fail: %v", err.Error())
	}

	var info UploadInfo
	err = mapstructure.Decode(n, &info)
	if err != nil {
		vErr := fmt.Errorf("storage uplaod file fail: %v", err.Error())
		return nil, vErr
	}
	info.FileName = fileName
	return &info, nil
}

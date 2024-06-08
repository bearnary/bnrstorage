package bnrstorage

import (
	"context"
)

type Bucket interface {
	UploadFile(ctx context.Context, fileName, path, storagePath, contentType string) (*UploadInfo, error)
	GetObjectUrl(ctx context.Context, fileName string) (string, error)
	GetObjectDownloadUrl(ctx context.Context, fileName string) (string, error)
	DeleteObject(ctx context.Context, fileName string) error
}

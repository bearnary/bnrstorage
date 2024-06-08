package bnrstorage

import (
	"context"

	"github.com/minio/minio-go/v7"
)

type defaultBucket struct {
	bucketName         string
	region             string
	mc                 *minio.Client
	ctx                context.Context
	cloudFlareProxyUrl *string
}

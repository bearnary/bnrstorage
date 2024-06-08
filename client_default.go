package bnrstorage

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type defaultClient struct {
	mc     *minio.Client
	region string
}

func NewClient(ctx context.Context, url, accessKeyId, secretAccessKey, region string, isSecure bool) (Client, error) {
	mc, err := minio.New(url, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyId, secretAccessKey, ""),
		Secure: isSecure,
	})
	if err != nil {
		return nil, fmt.Errorf("storage internal error : %v", err.Error())
	}
	return &defaultClient{
		mc:     mc,
		region: region,
	}, nil
}

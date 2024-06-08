package bnrstorage

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type defaultClient struct {
	ctx    context.Context
	cfg    *StorageConfig
	mc     *minio.Client
	region string
}

func NewClient(ctx context.Context, cfg *StorageConfig) (Client, error) {
	mc, err := minio.New(cfg.URL, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKeyId, cfg.SecretAccessKey, ""),
		Secure: cfg.IsSecure,
	})
	if err != nil {
		return nil, fmt.Errorf("storage internal error : %v", err.Error())
	}
	return &defaultClient{
		ctx:    ctx,
		cfg:    cfg,
		mc:     mc,
		region: cfg.Region,
	}, nil
}

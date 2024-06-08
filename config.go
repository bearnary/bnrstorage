package bnrstorage

type StorageConfig struct {
	URL               string `env:"STORAGE_URL"`
	AccessKeyId       string `env:"STORAGE_ACCESS_KEY_ID"`
	SecretAccessKey   string `env:"STORAGE_SECRET_ACCESS_KEY"`
	Region            string `env:"STORAGE_REGION"`
	IsSecure          bool   `env:"STORAGE_IS_SECURE"`
	DefaultBucketName string `env:"STORAGE_DEFAULT_BUCKET_NAME"`
}

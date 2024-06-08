package bnrstorage

type Client interface {
	SelectDefaultBucket() (Bucket, error)
	SelectBucket(bucketName string) (Bucket, error)
}

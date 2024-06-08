package bnrstorage

func (c *defaultClient) SelectDefaultBucket() (Bucket, error) {
	return c.SelectBucket(c.cfg.DefaultBucketName)
}

package bnrstorage

import "time"

type UploadInfo struct {
	Bucket       string
	FileName     string
	FullPath     string `mapstructure:"Key"`
	ETag         string
	Size         int64
	LastModified time.Time
	Location     string
	VersionID    string

	// Lifecycle expiry-date and ruleID associated with the expiry
	// not to be confused with `Expires` HTTP header.
	Expiration       time.Time
	ExpirationRuleID string
}

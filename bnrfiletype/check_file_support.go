package bnrfiletype

import (
	"errors"
	"io/ioutil"

	"github.com/h2non/filetype"
)

func CheckFileSupport(fullPath string, contentType string) error {

	buf, _ := ioutil.ReadFile(fullPath)
	kind, _ := filetype.Match(buf)
	if kind == filetype.Unknown {
		return errors.New("error un-support file type: unknown")
	}

	ft, err := GetFileTypeByContentType(contentType)
	if err != nil {
		return err
	}

	if filetype.IsImage(buf) && ft == Image {
		return nil
	}
	return errors.New("error un-support file type: non-image")
}

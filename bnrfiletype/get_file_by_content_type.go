package bnrfiletype

import (
	"errors"
)

func GetFileTypeByContentType(ct string) (FileType, error) {

	switch ct {
	case "image/jpeg", "image/png":
		return Image, nil
	}
	return Unsupported, errors.New("error unsupport file type: non image")
}

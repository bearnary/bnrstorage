package bnrfiletype

type FileType string

const (
	Image       FileType = "image"
	Unsupported FileType = ""
)

func (ft FileType) String() string {
	return string(ft)
}

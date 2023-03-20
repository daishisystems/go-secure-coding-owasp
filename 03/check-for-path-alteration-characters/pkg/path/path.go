package path

import (
	"path"
)

type Path struct{}

func NewPath() *Path {
	return &Path{}
}

func (*Path) PathIsValid(uri string) bool {
	cleanPath := path.Clean(uri)
	return uri == cleanPath
}

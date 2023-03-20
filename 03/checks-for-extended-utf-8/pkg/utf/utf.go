package utf

import (
	"unicode/utf8"
)

type Utf struct{}

func NewUtf() *Utf {
	return &Utf{}
}

func (*Utf) InputIsValidUtf8(input string) bool {
	return utf8.ValidString(input)
}

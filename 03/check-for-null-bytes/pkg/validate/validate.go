package validate

import (
	"strings"
)

type Validator struct{}

func NewNullByteCheck() *Validator {
	return &Validator{}
}

func (iv *Validator) ContainsNullByte(input string) bool {
	if input == "" {
		return false
	}
	return strings.Contains(input, "\x00")
}

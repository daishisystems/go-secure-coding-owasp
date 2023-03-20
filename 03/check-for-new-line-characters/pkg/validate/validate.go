package validate

import "strings"

type Validator struct{}

func NewValidator() *Validator {
	return &Validator{}
}

func (*Validator) ContainsNewLine(input string) bool {
	index := strings.IndexByte(input, '\n')
	return index >= 0
}

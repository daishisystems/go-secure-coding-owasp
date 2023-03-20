package validate

type Validator struct{}

func NewValidator() *Validator {
	return &Validator{}
}

func (*Validator) ContainsNewLine(input string) bool {
	for _, char := range input {
		if char == '\n' {
			return true
		}
	}
	return false
}

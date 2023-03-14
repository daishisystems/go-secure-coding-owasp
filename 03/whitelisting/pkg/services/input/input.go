package input

type InputService interface {
	CheckInput(input string) bool
}

type InputValues struct {
	allowedValues map[string]bool
}

func NewInputValues() *InputValues {
	return &InputValues{
		allowedValues: map[string]bool{
			"foo": true,
			"bar": true,
			"baz": true,
		},
	}
}

func (iv *InputValues) CheckInput(input string) bool {
	_, ok := iv.allowedValues[input]
	return ok
}

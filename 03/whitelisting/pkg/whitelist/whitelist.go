package whitelist

type Whitelist struct {
	allowedValues map[string]bool
}

func NewWhitelist() *Whitelist {
	return &Whitelist{
		allowedValues: map[string]bool{
			"foo": true,
			"bar": true,
		},
	}
}

func (iv *Whitelist) Check(input string) bool {
	_, ok := iv.allowedValues[input]
	return ok
}

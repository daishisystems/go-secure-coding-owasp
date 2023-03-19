package number

import (
	"fmt"
	"strconv"
)

type Number struct{}

func NewNumber() *Number {
	return &Number{}
}

func (iv *Number) IsNumber(input string) (bool, error) {
	_, err := strconv.Atoi(input)

	if err != nil {
		return false, fmt.Errorf("input '%s' is not a number", input)
	}

	return true, nil
}

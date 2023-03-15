package validator // todo: Separate interface from implementation?

import (
	"fmt"
	"strconv"
)

type IntegerValidator interface {
	IsValidInteger(num1Str, num2Str string) (int, int, error)
}

type IntegerValidate struct{}

func (s *IntegerValidate) IsValidInteger(num1Str, num2Str string) (int, int, error) {
	num1, err1 := strconv.Atoi(num1Str)
	num2, err2 := strconv.Atoi(num2Str)
	if err1 != nil || err2 != nil {
		return 0, 0, fmt.Errorf("invalid input")
	}
	if num1 < 1 || num1 > 100 || num2 < 1 || num2 > 100 {
		return 0, 0, fmt.Errorf("input out of range")
	}
	return num1, num2, nil
}

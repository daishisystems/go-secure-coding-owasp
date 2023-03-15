package calculator

type Calculator interface {
	Add(num1, num2 int) int
}

type Calculate struct{}

func (c *Calculate) Add(num1, num2 int) int {
	return num1 + num2
}

// func NewCalculator(calculator Calculator) *Calculate {
// 	return &Calculate{}
// }

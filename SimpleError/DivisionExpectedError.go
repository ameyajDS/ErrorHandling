package SimpleError

import "errors"

var ErrDivideByZero = errors.New("divide by zero")

func DivideExpectedError(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

package simpleerror

import (
	"errors"
	"fmt"
)

var ErrDivideByZero = errors.New("divide by zero")

func DivideExpectedError(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

func mainOne() {
	oprd1, oprd2 := 10, 0
	result, err := DivideExpectedError(oprd1, oprd2)
	if err != nil {
		switch {
		case errors.Is(err, ErrDivideByZero):
			fmt.Println("divide by zero error")
		default:
			fmt.Printf("unexpected division error: %s\n", err)
		}
		return
	}

	fmt.Printf("%d / %d = %d\n", oprd1, oprd2, result)

	//OUTPUT: divide by zero error
}

package ErrorHandling

import (
	"ErrorHandling/CustomError"
	"ErrorHandling/SimpleError"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	oprd1, oprd2 := 10, 0
	result, err := SimpleError.DivideExpectedError(oprd1, oprd2)
	if err != nil {
		switch {
		case errors.Is(err, SimpleError.ErrDivideByZero):
			fmt.Println("divide by zero error")
		default:
			fmt.Printf("unexpected division error: %s\n", err)
		}
		return
	}

	fmt.Printf("%d / %d = %d\n", oprd1, oprd2, result)
}

func mainOne() {
	oprd1, oprd2 := 10, 0
	result, err := SimpleError.Divide(oprd1, oprd2)
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%d / %d = %d\n", oprd1, oprd2, result)
}

func mainTwo() {
	ctx := context.Background().(*gin.Context)
	CustomError.ValidateTransfer(ctx)
}

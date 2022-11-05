package ErrorHandling

import (
	"ErrorHandling/customerror"
	"ErrorHandling/simpleerror"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

func mainOne() {
	oprd1, oprd2 := 10, 0
	result, err := simpleerror.Divide(oprd1, oprd2)
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%d / %d = %d\n", oprd1, oprd2, result)
}

func main() {
	oprd1, oprd2 := 10, 0
	result, err := simpleerror.DivideExpectedError(oprd1, oprd2)
	if err != nil {
		switch {
		case errors.Is(err, simpleerror.ErrDivideByZero):
			fmt.Println("divide by zero error")
		default:
			fmt.Printf("unexpected division error: %s\n", err)
		}
		return
	}

	fmt.Printf("%d / %d = %d\n", oprd1, oprd2, result)
}

func mainTwo() {
	ctx := context.Background().(*gin.Context)
	customerror.ValidateTransfer(ctx)
}

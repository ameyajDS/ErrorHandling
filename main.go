package ErrorHandling

import (
	"ErrorHandling/customerror"
	"context"
	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background().(*gin.Context)
	customerror.ValidateTransfer(ctx)
}

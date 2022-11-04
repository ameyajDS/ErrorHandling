package CustomError

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ValidateTransfer(context *gin.Context) {

	var request TransferRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, ErrMissingParameter)
	}
	if checkValidAccountNumber(request.PayerAcc) {
		context.AbortWithStatusJSON(http.StatusBadRequest,
			NewApiWithAdditionalData(
				ErrMissingParamCode,
				ErrMissingParamMessage,
				http.StatusBadRequest,
				map[string]interface{}{"Invalid payer account number": request.PayerAcc}))
	}
	if checkValidAccountNumber(request.PayeeAcc) {
		context.AbortWithStatusJSON(http.StatusBadRequest,
			NewApiWithAdditionalData(
				ErrMissingParamCode,
				ErrMissingParamMessage,
				http.StatusBadRequest,
				map[string]interface{}{"Invalid payee account number": request.PayeeAcc}))
	}

	context.JSON(http.StatusOK, nil)
}

func checkValidAccountNumber(accountNumber string) bool {
	if len(accountNumber) != 8 {
		return false
	}
	return true
}

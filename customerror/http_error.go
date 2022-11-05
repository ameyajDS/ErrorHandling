package customerror

import "net/http"

func New(errorCode string, errorMessage string, httpStatusCode int) *APIError {
	return &APIError{
		ErrorResponse: ErrorResponse{
			Code:    errorCode,
			Message: errorMessage,
		},
		HttpStatusCode: httpStatusCode,
	}
}

func NewApiWithAdditionalData(errorCode string, errorMessage string, httpStatusCode int, additionalData map[string]interface{}) *APIError {
	return &APIError{
		HttpStatusCode: httpStatusCode,
		ErrorResponse: ErrorResponse{
			Code:           errorCode,
			Message:        errorMessage,
			AdditionalData: additionalData,
		},
	}
}

type APIError struct {
	HttpStatusCode int
	ErrorResponse  ErrorResponse
}

// ERROR CODE
const (
	ErrAccountNumberValidationFailedCode = "ERR_ACCOUNT_NUMBER_INVALID"
	ErrMissingParamCode                  = "ERR_MISSING_PARAMETER_IN_REQUEST"
)

// ERROR MESSAGE
const (
	ErrInvalidAccNumberMessage = "Unable to complete the operation due to invalid account number"
	ErrMissingParamMessage     = "required mandatory request params are missing"
)

var (
	ErrMissingParameter = New(ErrMissingParamCode, ErrMissingParamMessage, http.StatusBadRequest)
)

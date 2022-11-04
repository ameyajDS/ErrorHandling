package CustomError

type ErrorResponse struct {
	Code           string                 `json:"errorCode"`
	Message        string                 `json:"errorMessage"`
	AdditionalData map[string]interface{} `json:"additionalData,omitempty"`
}

func NewErrorResponse(code string, message string, additionalData map[string]interface{}) *ErrorResponse {
	return &ErrorResponse{
		Code:           code,
		Message:        message,
		AdditionalData: additionalData,
	}
}

func (errorResponse *ErrorResponse) Error() string { return errorResponse.Message }

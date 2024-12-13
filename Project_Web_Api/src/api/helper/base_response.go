package helper

import "github.com/javadkavossi/GoLang_learning/api/validations"

type BaseHttpResponse struct {
	Result          any                            `json:"result"`
	Success         bool                           `json:"success"`
	ResultCode      int                            `json:"resultCode"`
	ValidationError *[]validations.ValidationError `json:"validationErrors"`
	Error           any                            `json:"error"`
}

type Result struct {
}

type SuccessResponse struct {
	Result          *Result `json:"result" `
	ResultCode      int     `json:"resultCode" default:"200"`
	Success         bool    `json:"success" default:"true"`
	ValidationError string  `json:"validationErrors" default:"null"`
	error           string  `json:"error" default:"null"`
}

type FailureResponse struct {
	Error           string                         `json:"error"`
	Result          *Result                        `json:"result"`
	ResultCode      int                            `json:"resultCode"`
	Success         bool                           `json:"success" default:"false"`
	ValidationError *[]validations.ValidationError `json:"validationErrors"`
}

func GenerateBaseResponse(result any, success bool, resultCode int) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:          result,
		Success:         success,
		ResultCode:      resultCode,
		ValidationError: nil,
		Error:           nil,
	}
}

func GenerateBaseResponseWithError(result any, success bool, resultCode int, err error) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:          result,
		Success:         success,
		ResultCode:      resultCode,
		ValidationError: nil,
		Error:           err.Error(),
	}
}

func GenerateBaseResponseWithValidationError(result any, success bool, resultCode int, err error) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:     result,
		Success:    success,
		ResultCode: resultCode,
		Error:      validations.GetValidationErrors(err),
	}
}

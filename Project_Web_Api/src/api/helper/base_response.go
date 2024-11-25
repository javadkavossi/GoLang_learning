package helper

import "github.com/javadkavossi/GoLang_learning/api/validations"

type BaseHttpResponse struct {
	Result          any                            `json:"result"`
	Success         bool                           `json:"success"`
	ResultCode      int                            `json:"resultCode"`
	ValidationError *[]validations.ValidationError `json:"validationErrors"`
	Error           any                            `json:"error"`
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

package helper

import "github.com/javadkavossi/GoLang_learning/pkg/serviceErrors"

var StatusCodeMapping = map[string]int{
	//OTP
	serviceErrors.OtpExists:   409,
	serviceErrors.OtpUsed:     409,
	serviceErrors.OtpNotValid: 400,
}

func TranslateErrorToStatusCode(err error) int {
	value, ok := StatusCodeMapping[err.Error()]
	if !ok {
		return 500
	}
	return value
}

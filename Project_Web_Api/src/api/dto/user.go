package dto

type GetOtpRequest struct {
	MobilNumber string `json:"mobilNumber" binding:"required,mobile,min=11,max=11"`
}

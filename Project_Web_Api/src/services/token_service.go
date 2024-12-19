package services

import (
	"github.com/javadkavossi/GoLang_learning/config"
	"github.com/javadkavossi/GoLang_learning/pkg/logging"
)

type TokenService struct {
	logger *logging.Logger
	cfg    *config.Config
}


type tokenDto struct {
	UserId int
	FirstName string
	LastName string
	UserName string
	Email string
	Roles []string
	MobilNumber string

}



// func NewTokenService(cfg *config.Config)*UserService {

// }
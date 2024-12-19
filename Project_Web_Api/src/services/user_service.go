package services

import (
	"github.com/javadkavossi/GoLang_learning/api/dto"
	"github.com/javadkavossi/GoLang_learning/common"
	"github.com/javadkavossi/GoLang_learning/config"
	"github.com/javadkavossi/GoLang_learning/data/db"
	"github.com/javadkavossi/GoLang_learning/pkg/logging"
	"gorm.io/gorm"
)

type UserService struct {
	logger     logging.Logger
	cfg        *config.Config
	otpService *OtpService
	database   *gorm.DB
}

func NewUserService(cfg *config.Config) *UserService {
	database := db.GetDb()
	logging := logging.NewLogger(cfg)

	return &UserService{
		cfg:        cfg,
		database:   database,
		logger:     logging,
		otpService: NewOtpService(cfg),
	}

}

func (s *UserService) SendOtp(req *dto.GetOtpRequest) error {
	otp := common.GenerateOtp()
	err := s.otpService.SetOtp(req.MobilNumber, otp)
	if err != nil {
		return err
	}
	return nil
}

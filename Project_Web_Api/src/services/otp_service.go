package services

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/javadkavossi/GoLang_learning/config"
	"github.com/javadkavossi/GoLang_learning/constants"
	"github.com/javadkavossi/GoLang_learning/data/cache"
	"github.com/javadkavossi/GoLang_learning/pkg/logging"
	"github.com/javadkavossi/GoLang_learning/pkg/serviceErrors"
)

type OtpService struct {
	logger      *logging.Logger
	cfg         *config.Config
	redisClient *redis.Client
}

type OtpDto struct {
	Value string
	Used  bool
}

func NewOtpService(cfg *config.Config) *OtpService {
	logger := logging.NewLogger(cfg)
	redisClient := cache.GetRedis()
	return &OtpService{logger: &logger, cfg: cfg, redisClient: redisClient}
}

func (s *OtpService) SetOtp(mobilNumber string, otp string) error {
	key := fmt.Sprintf("%s:%s", constants.RedisOtpDefaultKey, mobilNumber)
	val := &OtpDto{
		Value: otp,
		Used:  false,
	}
	res, err := cache.Get[OtpDto](s.redisClient, key)
	if err == nil && !res.Used {
		return &serviceErrors.ServiceError{EndUserMessage: serviceErrors.OtpExists}
	} else if err == nil && res.Used {
		return &serviceErrors.ServiceError{EndUserMessage: serviceErrors.OtpUsed}

	}
	err = cache.Set[*OtpDto](s.redisClient, key, val, s.cfg.Otp.ExpireTime*time.Second)
	if err != nil {
		return err
	}
	return nil

}

func (s *OtpService) VerifyOtp(mobilNumber string, otp string) error {
	key := fmt.Sprintf("%s:%s", constants.RedisOtpDefaultKey, mobilNumber)
	res, err := cache.Get[OtpDto](s.redisClient, key)
	if err != nil {
		return err
	} else if err == nil && res.Used {
		return &serviceErrors.ServiceError{EndUserMessage: serviceErrors.OtpUsed}
	} else if err == nil && !res.Used && res.Value != otp {
		return &serviceErrors.ServiceError{EndUserMessage: serviceErrors.OtpNotValid}
	} else if err == nil && !res.Used && res.Value == otp {
		res.Used = true
		err = cache.Set(s.redisClient, key, res, s.cfg.Otp.ExpireTime*time.Second)
	}
	return nil

}

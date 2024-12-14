package logging

import "github.com/javadkavossi/GoLang_learning/config"

type Logger interface {
	Init()

	Debug(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{})
	Debugf(template string, args ...interface{})

	Info(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{})
	Infof(template string, args ...interface{})

	Warn(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{})
	Warnf(template string, args ...interface{})

	Error(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{})
	Errorf(template string, args ...interface{})

	Fatal(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{})
	Fatalf(template string, args ...interface{})
}

func NewLogger(cfg *config.Config) Logger {
	if cfg.Logger.Logger == "zap" {
		return newZapLogger(cfg)
	} else if cfg.Logger.Logger == "zerolog" {
		return newZeroLogger(cfg)
	}
	panic("logger not support")
}

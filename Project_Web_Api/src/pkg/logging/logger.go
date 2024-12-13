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

	Error(err error, category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{})
	Errorf(err error, template string, args ...interface{})

	Fatal(err error, category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{})
	Fatalf(err error, template string, args ...interface{})
}

func NewLogger(cfg config.Config) Logger {
	return NewZapLogger(cfg)
}
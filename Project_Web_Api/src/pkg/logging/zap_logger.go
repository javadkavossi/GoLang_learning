package logging

import (
	"github.com/javadkavossi/GoLang_learning/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logLevelMap = map[string]zapcore.Level{

	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"fatal": zapcore.FatalLevel,
}

type zapLogger struct {
	cfg    *config.Config
	logger *zap.SugaredLogger
}

func NewZapLogger(cfg *config.Config) *zapLogger {
	logger := &zapLogger{cfg: cfg}
	logger.Init()
	return logger

}

func (logger *zapLogger) Init() {
	writer :=zapcore.AddSync(&lumberjack.Logger{
		Filename: logger.cfg.Logger.FilePath,
		MaxSize:10,
		MaxAge:5,
		LocalTime:true,
		MaxBackups:10,
		Compress:true,


	})
}

func Debug(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{}) {

}

func Debugf(template string, args ...interface{}) {

}

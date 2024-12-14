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

func newZapLogger(cfg *config.Config) *zapLogger {
	logger := &zapLogger{cfg: cfg}
	logger.Init()
	return logger

}
func (l *zapLogger) getLogLevel() zapcore.Level {
	level, exists := logLevelMap[l.cfg.Logger.Level]
	if !exists {
		return zap.DebugLevel
	}
	return level
}

func (l *zapLogger) Init() {
	writer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   l.cfg.Logger.FilePath,
		MaxSize:    10,
		MaxAge:     5,
		LocalTime:  true,
		MaxBackups: 10,
		Compress:   true,
	})

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		writer,
		l.getLogLevel(),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel)).Sugar().
		With("AppName", "DPGGameService", "LoggerName", "Zerolog")

	l.logger = logger

}

// ========================================== Debug =====================================================

func (l *zapLogger) Debug(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(extra, category, subCategory)
	l.logger.Debugw(message, params...)
}
func (l *zapLogger) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args)

}

// ========================================== Info  =====================================================

func (l *zapLogger) Info(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(extra, category, subCategory)
	l.logger.Infow(message, params...)
}
func (l *zapLogger) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args)
}

// ========================================== Warning =====================================================

func (l *zapLogger) Warn(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(extra, category, subCategory)
	l.logger.Warnw(message, params...)
}
func (l *zapLogger) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args)
}

// ========================================== Error =====================================================

func (l *zapLogger) Error(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(extra, category, subCategory)
	l.logger.Errorw(message, params...)
}
func (l *zapLogger) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args)
}

// ========================================== Fatal =====================================================

func (l *zapLogger) Fatal(category Category, subCategory SubCategory, message string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(extra, category, subCategory)
	l.logger.Fatalw(message, params...)
}
func (l *zapLogger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args)
}

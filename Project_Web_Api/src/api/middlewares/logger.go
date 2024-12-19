package middlewares

import (
	"bytes"
	"io/ioutil"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/javadkavossi/GoLang_learning/config"
	"github.com/javadkavossi/GoLang_learning/pkg/logging"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)

}

func DefaultStructuredLogger(cfg *config.Config) gin.HandlerFunc {
	logger := logging.NewLogger(cfg)
	return structuredLogger(logger)
}

func structuredLogger(logger logging.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if strings.Contains(ctx.FullPath(), "swagger") {
			ctx.Next()
		} else {
			blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
			start := time.Now() // Start
			path := ctx.FullPath()
			raw := ctx.Request.URL.RawQuery
			bodyBytes, _ := ioutil.ReadAll(ctx.Request.Body)
			ctx.Request.Body.Close()
			ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

			ctx.Writer = blw
			ctx.Next()

			param := gin.LogFormatterParams{}

			param.TimeStamp = time.Now() // Stop
			param.Latency = param.TimeStamp.Sub(start)
			param.ClientIP = ctx.ClientIP()
			param.Method = ctx.Request.Method
			param.StatusCode = ctx.Writer.Status()
			param.ErrorMessage = ctx.Errors.ByType(gin.ErrorTypePrivate).String()
			param.BodySize = ctx.Writer.Size()
			if raw != "" {
				path = raw
			}
			param.Path = path

			keys := map[logging.ExtraKey]interface{}{}
			keys[logging.Path] = param.Path
			keys[logging.ClientIp] = param.ClientIP
			keys[logging.Method] = param.Method
			keys[logging.Latency] = param.Latency
			keys[logging.StatusCode] = param.StatusCode
			keys[logging.ErrorMessage] = param.ErrorMessage
			keys[logging.BodySize] = param.BodySize
			keys[logging.RequestBody] = string(bodyBytes)
			keys[logging.ResponseBody] = blw.body.String()

			logger.Info(logging.RequestResponse, logging.Api, "", keys)
		}
	}

}

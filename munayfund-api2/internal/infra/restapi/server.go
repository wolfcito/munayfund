package restapi

import (
	"fmt"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func NewServer(appName string) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	server := gin.New()

	server.Use(
		requestid.New(),
		gin.Recovery(),
		gin.LoggerWithConfig(GetLoggerConfig(appName)),
	)

	return server
}

func GetLoggerConfig(appName string) gin.LoggerConfig {
	var formatter = func(param gin.LogFormatterParams) string {
		if param.Latency > time.Minute {
			param.Latency = param.Latency - param.Latency%time.Second
		}

		return fmt.Sprintf(
			"[method:%s][time:%s][status:%d][latency:%s][client_ip:%s][endpoint:%s]\n",
			param.Method,
			param.TimeStamp.Format("2006/01/02 - 15:04:05"),
			param.StatusCode,
			param.Latency,
			param.ClientIP,
			param.Request.URL,
		)
	}

	return gin.LoggerConfig{
		Formatter: formatter,
		SkipPaths: []string{
			fmt.Sprintf("/api/%s/health-check", appName),
		},
	}
}

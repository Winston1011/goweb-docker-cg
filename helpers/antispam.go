package helpers

import (
	"github.com/gin-gonic/gin"
	"goweb-docker-cg/pkg/antispam"
	"goweb-docker-cg/pkg/antispam/logger"
	"goweb-docker-cg/pkg/golib/v2/zlog"
)

func initAntiSpam() {
	antispam.InitAntiSpam(setAntiSpamLogger)
}

// 使用zlog重写zns的logger
type antiLogger struct{}

func (l *antiLogger) Logger(ctx *gin.Context, level logger.LogLevel, msg string) {
	switch level {
	case logger.Debug:
		zlog.Debug(ctx, msg)
	case logger.Info:
		zlog.Info(ctx, msg)
	case logger.Warn:
		zlog.Warn(ctx, msg)
	case logger.Error:
		zlog.Error(ctx, msg)
	}
}

func setAntiSpamLogger(conf *antispam.Config) {
	conf.Logger = &antiLogger{}
}

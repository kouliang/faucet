package route

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"faucet-app/setting"
)

// ginLogger 接收gin框架默认的日志
func ginLogger(mode string, cfg *setting.LogConfig) gin.HandlerFunc {

	//encoder
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	//WriteSyncer
	lumberJackLogger := &lumberjack.Logger{
		Filename:   cfg.FileName,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   false,
	}
	writeSyncer := zapcore.AddSync(lumberJackLogger)

	//logLevel
	logLevel := new(zapcore.Level)
	logLevel.UnmarshalText([]byte(cfg.Level))

	var core zapcore.Core
	if mode == "dev" {
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, logLevel),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewCore(encoder, writeSyncer, logLevel)
	}
	zap.ReplaceGlobals(zap.New(core, zap.AddCaller())) //替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()就能找到logger

	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.Request.URL.Path
		query := ctx.Request.URL.RawQuery
		ctx.Next()

		cost := time.Since(start)
		zap.L().Info(path,
			zap.Int("status", ctx.Writer.Status()),
			zap.String("method", ctx.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", ctx.ClientIP()),
			zap.String("user-agent", ctx.Request.UserAgent()),
			zap.String("errors", ctx.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

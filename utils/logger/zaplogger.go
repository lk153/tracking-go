package logger

import (
	"log"
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//ZapLogger ...
var ZapLogger *zap.Logger

func init() {
	cfg := zap.NewDevelopmentConfig()
	if os.Getenv("PROD") == "yes" {
		cfg = zap.NewProductionConfig()
	} else {
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	var e error
	ZapLogger, e = cfg.Build()
	if e != nil {
		log.Fatalf("Unable to construct logger %s", e.Error())
	}

}

//SyncLog ...
func SyncLog() {
	_ = ZapLogger.Sync()
}

//Logger ...
type Logger = logr.Logger

// logr level is invert of zap log level
// see more https://github.com/go-logr/zapr#implementation-details
const (
	LogDebugLevel int = -int(zapcore.DebugLevel)
	LogInfoLevel      = -int(zapcore.InfoLevel)
	LogWarnLevel      = -int(zapcore.WarnLevel)
	LogErrorLevel     = -int(zapcore.ErrorLevel)
	LogFatalLevel     = -int(zapcore.FatalLevel)
)

//GetLoggerFactory ...
func GetLoggerFactory(name string) Logger {
	return zapr.NewLogger(ZapLogger).WithName(name)
}

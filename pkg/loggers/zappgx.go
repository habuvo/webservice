package loggers

import (
	"context"

	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	logger *zap.Logger
}

func NewZapLogger(logger *zap.Logger) *ZapLogger {
	return &ZapLogger{logger: logger.WithOptions(zap.AddCallerSkip(1))}
}

func (l *ZapLogger) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	fields := make([]zapcore.Field, len(data))
	i := 0
	for k, v := range data {
		fields[i] = zap.Any(k, v)
		i++
	}

	switch level {
	case pgx.LogLevelTrace:
		l.logger.Debug(msg, append(fields, zap.Stringer("PGX_LOG_LEVEL", level))...)
	case pgx.LogLevelDebug:
		l.logger.Debug(msg, fields...)
	case pgx.LogLevelInfo:
		l.logger.Info(msg, fields...)
	case pgx.LogLevelWarn:
		l.logger.Warn(msg, fields...)
	case pgx.LogLevelError:
		l.logger.Error(msg, fields...)
	default:
		l.logger.Error(msg, append(fields, zap.Stringer("PGX_LOG_LEVEL", level))...)
	}
}

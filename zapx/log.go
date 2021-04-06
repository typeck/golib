package zapx

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// DefaultLumberJackLogger default
func DefaultLumberJackLogger() *lumberjack.Logger{
	return 	&lumberjack.Logger{
		MaxSize:    500,
		MaxBackups: 10,
		MaxAge:     1,
		Compress:   false,
	}
}

var (
	//Levels log level
	Levels = []zapcore.Level{zapcore.DebugLevel, zapcore.InfoLevel, zapcore.WarnLevel, zapcore.ErrorLevel, zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel}
)

func wrapFileName(logger *lumberjack.Logger, level zapcore.Level) *lumberjack.Logger {
	fileName := level.String() + ".log"
	if logger.Filename == "" {
		logger.Filename = fileName
		return logger
	}
	logger.Filename = logger.Filename + "-" + fileName
	return logger
}

// NewLumberJackLoggerFunc  writer
type NewLumberJackLoggerFunc func() *lumberjack.Logger

//ProductionTee .
func ProductionTee(loggerFunc NewLumberJackLoggerFunc) zapcore.Core {
	if loggerFunc == nil {
		loggerFunc = DefaultLumberJackLogger
	}
	var cores  []zapcore.Core
	for _, l := range Levels {
		core := zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			zapcore.AddSync(wrapFileName(loggerFunc(), l)), SeparateLevel(l))
		cores = append(cores, core)
	}
	core := zapcore.NewTee(cores...)
	return core
}

//DevelopmentTee .
func DevelopmentTee(loggerFunc NewLumberJackLoggerFunc) zapcore.Core {
	if loggerFunc == nil {
		loggerFunc = DefaultLumberJackLogger
	}
	var cores  []zapcore.Core
	for _, l := range Levels {
		core := zapcore.NewCore(zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
			zapcore.AddSync(wrapFileName(loggerFunc(), l)), SeparateLevel(l))
		cores = append(cores, core)
	}
	core := zapcore.NewTee(cores...)
	return core
}

//SeparateLevel .
func SeparateLevel(l zapcore.Level) zapcore.LevelEnabler {
	return  zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == l
	})
}
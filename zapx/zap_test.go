package zapx

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"testing"
)

func TestProductionTee(t *testing.T) {
	logger := zap.New(ProductionTee(func() *lumberjack.Logger {
		l := DefaultLumberJackLogger()
		l.Filename = "./zapx"
		return l
	}))
	logger.Info("------------------------")
	logger.Error("+++++++++++++++++++++++")

	zap.ReplaceGlobals(logger)
	zap.L().Info("example")
}

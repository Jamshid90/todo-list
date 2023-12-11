package logger

import (
	"fmt"

	"go.uber.org/zap"
)

const (
	EnvironmentProd = "production"
	EnvironmentDev  = "develop"
)

func productionConfig(file string) zap.Config {
	configZap := zap.NewProductionConfig()
	configZap.OutputPaths = []string{"stdout", file}
	configZap.DisableStacktrace = true
	return configZap
}

func developmentConfig(file string) zap.Config {
	configZap := zap.NewDevelopmentConfig()
	configZap.OutputPaths = []string{"stdout", file}
	configZap.ErrorOutputPaths = []string{"stderr"}
	return configZap
}

func setLevel(level string, cfg zap.Config) zap.Config {
	switch level {
	case "debug":
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		cfg.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		cfg.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	case "dpanic":
		cfg.Level = zap.NewAtomicLevelAt(zap.DPanicLevel)
	case "panic":
		cfg.Level = zap.NewAtomicLevelAt(zap.PanicLevel)
	case "fatal":
		cfg.Level = zap.NewAtomicLevelAt(zap.FatalLevel)
	default:
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}
	return cfg
}

func NewProduction(level, logFileName string) (*zap.Logger, error) {
	file := fmt.Sprintf("./%s.log", logFileName)
	configZap := productionConfig(file)
	configZap = setLevel(level, configZap)
	return configZap.Build()
}

func NewDevelopment(level, logFileName string) (*zap.Logger, error) {
	file := fmt.Sprintf("./%s.log", logFileName)
	configZap := developmentConfig(file)
	configZap = setLevel(level, configZap)
	return configZap.Build()
}

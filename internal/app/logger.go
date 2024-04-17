package app

import (
	"exchange-service/configuration"
	"log/slog"

	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/exp/zapslog"
	"go.uber.org/zap/zapcore"
)

var zapLogger *zap.Logger

func NewLogger(config configuration.Config) *slog.Logger {
	if zapLogger == nil {
		setupZapLogger(config)
	}
	return slog.New(zapslog.NewHandler(zapLogger.Core(), nil))
}

func setupZapLogger(config configuration.Config) {
	zapLogger = getLogger(config)
}

func getLogger(config configuration.Config) *zap.Logger {
	level := func() zap.AtomicLevel {
		switch config.Logging.Level {
		case configuration.Debug:
			return zap.NewAtomicLevelAt(zap.DebugLevel)
		case configuration.Info:
			return zap.NewAtomicLevelAt(zap.InfoLevel)
		case configuration.Warn:
			return zap.NewAtomicLevelAt(zap.WarnLevel)
		case configuration.Error:
			return zap.NewAtomicLevelAt(zap.ErrorLevel)
		default:
			panic("Unknown Logging Level")
		}

	}()
	encoderCfg := func() zapcore.EncoderConfig {
		switch config.App.Env {
		case configuration.Dev:
			return zap.NewDevelopmentEncoderConfig()
		default: // for Test And prod
			return zap.NewProductionEncoderConfig()
		}
	}()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	logConfig := zap.Config{
		Level:             level,
		Development:       true,
		DisableCaller:     true,
		DisableStacktrace: true,
		Sampling:          nil,
		Encoding:          "console",
		EncoderConfig:     encoderCfg,
		OutputPaths: []string{
			"stderr",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
	}
	return zap.Must(logConfig.Build())
}

type FxLogger struct {
	logger *slog.Logger
}

func NewFxLogger(logger *slog.Logger) FxLogger {
	return FxLogger{logger}
}

func (f FxLogger) LogEvent(event fxevent.Event) {
	logSilentInfo := func(eventName string, property string, value string) {
		// f.logger.Info(fmt.Sprintf("Fx event: %v", eventName), property, "silenced")
	}

	switch e := event.(type) {
	case *fxevent.OnStartExecuting:
		logSilentInfo("OnStartExecuting", "FunctionName", e.FunctionName)
	case *fxevent.OnStartExecuted:
		if f.logError(e.Err) == nil {
			logSilentInfo("OnStartExecuted", "FunctionName", e.FunctionName)
		}
	case *fxevent.OnStopExecuting:
		logSilentInfo("OnStopExecuting", "FunctionName", e.FunctionName)
	case *fxevent.OnStopExecuted:
		if f.logError(e.Err) == nil {
			logSilentInfo("OnStopExecuted", "FunctionName", e.FunctionName)
		}
	case *fxevent.Supplied:
		if f.logError(e.Err) == nil {
			logSilentInfo("Supplied", "TypeName", e.TypeName)
		}
	case *fxevent.Provided:
		if f.logError(e.Err) == nil {
			logSilentInfo("Provided", "ConstructorName", e.ConstructorName)
		}
	case *fxevent.Replaced:
		if f.logError(e.Err) == nil {
			logSilentInfo("Replaced", "ModuleName", e.ModuleName)
		}
	case *fxevent.Decorated:
		if f.logError(e.Err) == nil {
			logSilentInfo("Decorated", "DecoratorName", e.DecoratorName)
		}
	case *fxevent.Run:
		if f.logError(e.Err) == nil {
			logSilentInfo("Run", "Name", e.Name)
		}
	case *fxevent.Invoking:
		logSilentInfo("Invoking", "FunctionName", e.FunctionName)
	case *fxevent.Invoked:
		if f.logError(e.Err) == nil {
			logSilentInfo("Invoked", "FunctionName", e.FunctionName)
		}
	case *fxevent.Started:
		if f.logError(e.Err) == nil {
			f.logger.Info("Started")
		}
	case *fxevent.Stopping:
		logSilentInfo("Stopping", "Signal", e.Signal.String())
	case *fxevent.Stopped:
		if f.logError(e.Err) == nil {
			logSilentInfo("Stopped", "", "")
		}
	case *fxevent.RollingBack:
		logSilentInfo("RollingBack", "StartErr", e.StartErr.Error())
	case *fxevent.RolledBack:
		if f.logError(e.Err) == nil {
			logSilentInfo("RolledBack", "", "")
		}
	case *fxevent.LoggerInitialized:
		if f.logError(e.Err) == nil {
			logSilentInfo("LoggerInitialized", "ConstructorName", e.ConstructorName)
		}
	default:
		f.logger.Error("Unknown Event")
	}
}

func (f FxLogger) logError(e error) error {
	if e != nil {
		f.logger.Error(e.Error())
	}
	return e
}

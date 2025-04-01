package infrastructure

import (
	"itsm/config"
	"os"
	"runtime"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger holds the Zap logger instance
var Logger *zap.Logger

// InitLogger initializes the logger based on the provided configuration
func InitLogger(cfg *config.Config) {
	var zapConfig zap.Config

	// ✅ Ensure logs directory exists and is configurable
	logDir := cfg.Logging.Output
	if logDir != "stdout" {
		if _, err := os.Stat(logDir); os.IsNotExist(err) {
			if err := os.Mkdir(logDir, 0755); err != nil {
				panic("Failed to create logs directory: " + err.Error())
			}
		}
	}

	// ✅ Set logging format and level dynamically
	logLevel := zapcore.InfoLevel
	switch cfg.Logging.Level {
	case "debug":
		logLevel = zapcore.DebugLevel
	case "warn":
		logLevel = zapcore.WarnLevel
	case "error":
		logLevel = zapcore.ErrorLevel
	}

	if cfg.Logging.Format == "json" {
		zapConfig = zap.NewProductionConfig() // JSON logs (structured logging)
	} else {
		zapConfig = zap.NewDevelopmentConfig() // Console logs (for development)
		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	// ✅ Configure output paths for logs
	if logDir == "stdout" {
		zapConfig.OutputPaths = []string{"stdout"} // Console only
	} else {
		zapConfig.OutputPaths = []string{"stdout", logDir + "/service.log"} // Console + file
	}

	// ✅ Apply log level dynamically
	zapConfig.Level = zap.NewAtomicLevelAt(logLevel)

	// ✅ Build the logger
	var err error
	Logger, err = zapConfig.Build()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}

	// ✅ Ensure logger flushes before shutdown
	defer func() {
		err := Logger.Sync()
		if err != nil && runtime.GOOS == "windows" {
			// Ignore the sync error on Windows
			return
		}
		if err != nil {
			Logger.Error("Failed to sync logger", zap.Error(err))
		}
	}()

	Logger.Info("✅ Logger initialized successfully",
		zap.String("level", cfg.Logging.Level),
		zap.String("format", cfg.Logging.Format),
		zap.String("output", cfg.Logging.Output),
	)
}

// GetLogger returns the logger instance
func GetLogger() *zap.Logger {
	if Logger == nil {
		panic("🚨 Logger is not initialized! Call InitLogger() first.")
	}
	return Logger
}

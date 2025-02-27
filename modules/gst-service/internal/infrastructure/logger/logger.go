package logger

import (
	"gst-service/internal/infrastructure/config"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Loggers
var (
	AppLogger *zap.Logger    // Application logger using Zap
	SISLogger *logrus.Logger // Security (SIS) logger using Logrus
	logFile   *os.File       // File handle for logs
	once      sync.Once      // Ensures cleanup is only run once
)

// EnsureLogDirectoryExists checks and creates the logs directory if missing
func EnsureLogDirectoryExists(logFilePath string) error {
	logDir := "logs"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err := os.Mkdir(logDir, 0755); err != nil {
			return err
		}
	}
	return nil
}

// InitializeLoggers sets up both Zap (for app logs) and Logrus (for security logs)
func InitializeLoggers(cfg *config.Config) error {
	logFilePath := cfg.Logging.File // Read log file path from config

	// Ensure the logs directory exists
	if err := EnsureLogDirectoryExists(logFilePath); err != nil {
		return err
	}

	// Create a file writer for Zap logs
	var err error
	logFile, err = os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	// Zap Logger Configuration (Application Logging)
	zapConfig := zap.NewProductionEncoderConfig()
	zapConfig.TimeKey = "timestamp"
	zapConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zapConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(logFile), zapcore.AddSync(os.Stdout)), // ✅ Log to both file & console
		zap.InfoLevel,
	)
	AppLogger = zap.New(core)

	// Logrus Logger Configuration (SIS Security Logging)
	SISLogger = logrus.New()
	SISLogger.SetFormatter(&logrus.JSONFormatter{})
	SISLogger.SetOutput(logFile) // ✅ Log to file

	AppLogger.Info("✅ Application logger initialized")
	SISLogger.Info("✅ SIS security logger initialized")

	return nil
}

// CloseLoggers ensures logs are flushed before shutdown
func CloseLoggers() {
	once.Do(func() { // Ensures cleanup is executed only once
		if AppLogger != nil {
			_ = AppLogger.Sync() // ✅ Force Zap to flush logs
		}
		if logFile != nil {
			logFile.Sync() // ✅ Force OS to write logs to disk
			logFile.Close()
		}
	})
}

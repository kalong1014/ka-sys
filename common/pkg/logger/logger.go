package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// 初始化日志记录器
func InitLogger() *logrus.Logger {
	logger := logrus.New()

	// 设置输出格式为JSON
	logger.SetFormatter(&logrus.JSONFormatter{})

	// 设置输出到标准输出
	logger.SetOutput(os.Stdout)

	// 设置日志级别
	level, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		level = logrus.InfoLevel
	}
	logger.SetLevel(level)

	return logger
}

// 带上下文的日志记录器
type ContextLogger struct {
	logger  *logrus.Logger
	context map[string]interface{}
}

func NewContextLogger(logger *logrus.Logger) *ContextLogger {
	return &ContextLogger{
		logger:  logger,
		context: make(map[string]interface{}),
	}
}

func (l *ContextLogger) WithField(key string, value interface{}) *ContextLogger {
	newLogger := &ContextLogger{
		logger:  l.logger,
		context: make(map[string]interface{}),
	}

	// 复制现有上下文
	for k, v := range l.context {
		newLogger.context[k] = v
	}

	// 添加新字段
	newLogger.context[key] = value

	return newLogger
}

func (l *ContextLogger) Info(msg string) {
	l.logger.WithFields(l.context).Info(msg)
}

func (l *ContextLogger) Error(msg string, err error) {
	fields := logrus.Fields(l.context)
	fields["error"] = err.Error()
	l.logger.WithFields(fields).Error(msg)
}

// 其他日志级别方法...

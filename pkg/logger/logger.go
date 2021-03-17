package logger

import "go.uber.org/zap"

var (
	logger Logger
)

// Initialize default production is false if not call func
func Initialize(environment string) {
	var conf zap.Config
	conf = zap.NewProductionConfig()

	if environment != "production" {
		conf = zap.NewDevelopmentConfig()
	}

	conf.DisableStacktrace = true
	log, err := conf.Build()
	if err != nil {
		panic(err)
	}

	logger = log.WithOptions(zap.AddCallerSkip(1)).Sugar()
}

// Debug uses fmt.Sprint to construct and log a message
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// Debugf uses fmt.Sprintf to log a templated message
func Debugf(template string, args ...interface{}) {
	logger.Debugf(template, args...)
}

// Debugw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With
func Debugw(msg string, keysValues ...interface{}) {
	logger.Debugw(msg, keysValues...)
}

// Info uses fmt.Sprint to construct and log a message
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Infof uses fmt.Sprintf to log a templated message
func Infof(template string, args ...interface{}) {
	logger.Infof(template, args...)
}

// Infow logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Infow(msg string, keysValues ...interface{}) {
	logger.Infow(msg, keysValues...)
}

// Warn uses fmt.Sprint to construct and log a message
func Warn(args ...interface{}) {
	logger.Warn(args...)
}

// Warnf uses fmt.Sprintf to log a templated message
func Warnf(template string, args ...interface{}) {
	logger.Warnf(template, args...)
}

// Warnw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Warnw(msg string, keysValues ...interface{}) {
	logger.Warnw(msg, keysValues...)
}

// Error uses fmt.Sprint to construct and log a message
func Error(args ...interface{}) {
	logger.Error(args...)
}

// Errorf uses fmt.Sprintf to log a templated message
func Errorf(template string, args ...interface{}) {
	logger.Errorf(template, args...)
}

// Errorw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Errorw(msg string, keysValues ...interface{}) {
	logger.Errorw(msg, keysValues...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit
func Fatalf(template string, args ...interface{}) {
	logger.Fatalf(template, args...)
}

// Fatalw logs a message with some additional context, then calls os.Exit. The
// variadic key-value pairs are treated as they are in With
func Fatalw(msg string, keysValues ...interface{}) {
	logger.Fatalw(msg, keysValues...)
}

// Panic uses fmt.Sprint to construct and log a message, then panics
func Panic(args ...interface{}) {
	logger.Panic(args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics
func Panicf(template string, args ...interface{}) {
	logger.Panicf(template, args...)
}

// Panicw logs a message with some additional context, then panics. The
// variadic key-value pairs are treated as they are in With
func Panicw(msg string, keysValues ...interface{}) {
	logger.Panicw(msg, keysValues...)
}

// WithLogger set global logger by new logger
func WithLogger(_logger Logger) {
	logger = _logger
}

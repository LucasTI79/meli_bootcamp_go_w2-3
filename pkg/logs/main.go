package logs

import (
	"fmt"
	"net"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/joho/godotenv"
)

const (
	// InfoLevel level. Usually only used for production logs.
	InfoLevel = iota
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// ErrorLevel level. Used for errors that should definitely be noted.
	ErrorLevel
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. highest level of severity.
	FatalLevel
)

// levelNames maps the logging levels to their names.
var (
	levelNames = map[int]string{
		InfoLevel:  "INFO",
		DebugLevel: "DEBUG",
		WarnLevel:  "WARN",
		ErrorLevel: "ERROR",
		FatalLevel: "FATAL",
	}
)

// Logger is the interface used for logging in the application.
type Logger interface {
	// Info logs a message at the info log level.
	Info(msg string, args ...interface{})
	// Debug logs a message at the debug log level.
	Debug(msg string, args ...interface{})
	// Warn logs a message at the warn log level.
	Warn(msg string, args ...interface{})
	// Error logs a message at the error log level.
	Error(msg string, args ...interface{})
	// Fatal logs a message at the fatal log level.
	Fatal(msg string, args ...interface{})
	// WithFields adds a map of fields to the logger context.
	WithFields(fields map[string]interface{})
	// WithField adds a single field to the logger context.
	WithError(err error) Logger
	// SetLevel sets the logger level.
	SetLevel(level int)
	// GetLevel returns the logger level.
	GetLevel() int
	// Save logs to DB
	Save() error
	// SetLogAll sets the logger to log all messages regardless of the log level.
	SetLogAll(logAll bool)
	// Printf logs a message using a format string and arguments.
	Printf(format string, args ...interface{})
	// SetLevelColor sets the color for a log level.
	SetLevelColor(level int, color string)
	// log logs a message at the specified log level.
	log(level int, msg string, args ...interface{})
}

type logger struct {
	level       int
	fields      map[string]interface{}
	logAll      bool
	levelColors map[int]string
}

// NewLogger creates a new logger instance.
// The logger level can be set with the environment variable DEBUG.
// If DEBUG is set to true, all messages will be logged regardless of the log level.
func NewLogger() Logger {
	_ = godotenv.Load(".env")
	isDebugMode := strings.ToLower(os.Getenv("DEBUG")) == "true"
	if isDebugMode {
		fmt.Println("[LOGZILLA DETAILS][WARNING] DEBUG MODE IS ON ITS NOT RECOMMENDED FOR PRODUCTION ENVIRONMENTS [WARNING")
	}

	return &logger{
		levelColors: map[int]string{
			InfoLevel:  "\033[32m", // green
			DebugLevel: "\033[36m", // cyan
			WarnLevel:  "\033[33m", // yellow
			ErrorLevel: "\033[31m", // red
			FatalLevel: "\033[35m", // magenta
		},
		logAll: isDebugMode,
	}
}

// ShowInfo shows the logger info.
func (l *logger) ShowInfo() {
	l.Info("[LOGZILLA DETAILS] [INFO] Logger Info")
	l.Info("[LOGZILLA DETAILS] [INFO] Level: ", l.level)
	l.Info("[LOGZILLA DETAILS] [INFO] Fields: ", l.fields)
	l.Info("[LOGZILLA DETAILS] [INFO] LogAll: ", l.logAll)
	l.Info("[LOGZILLA DETAILS] [INFO] LevelColors: ", l.levelColors)
}

// SetLevel sets the logger level.
func (l *logger) SetLevel(level int) {
	l.level = level
}

// SetLevelColor sets the color for a log level.
func (l *logger) SetLevelColor(level int, color string) {
	l.levelColors[level] = color

}

// GetLevel returns the logger level.
func (l *logger) GetLevel() int {
	return l.level
}

// WithFields adds a map of fields to the logger context.
func (l *logger) WithFields(fields map[string]interface{}) {
	l.fields = fields
}

// WithError adds an error to the logger context.
func (l *logger) WithError(err error) Logger {
	return l
}

// Fatal logs a message at the fatal log level.
func (l *logger) Fatal(msg string, args ...interface{}) {
	if l.level >= FatalLevel {
		l.log(FatalLevel, msg, args...)
		os.Exit(1)
	}
}

// Error logs a message at the error log level.
func (l *logger) Error(msg string, args ...interface{}) {
	if l.level >= ErrorLevel {
		l.log(ErrorLevel, msg)
	}
}

// Warn logs a message at the warn log level.
func (l *logger) Warn(msg string, args ...interface{}) {
	if l.level >= WarnLevel {
		l.log(WarnLevel, msg)
	}
}

// Debug logs a message at the debug log level.
func (l *logger) Debug(msg string, args ...interface{}) {
	if l.level >= DebugLevel {
		l.log(DebugLevel, msg)
	}
}

// Info logs a message at the info log level.
func (l *logger) Info(msg string, args ...interface{}) {
	if l.level >= InfoLevel {
		l.log(InfoLevel, msg)
	}
}

// log logs a message at the specified log level.
func (l *logger) log(level int, msg string, args ...interface{}) {
	if l.logAll || l.level >= level && (isLocalHost() || level >= WarnLevel) {
		_, file, line, _ := runtime.Caller(2)
		fileName := path.Base(file)
		msg = fmt.Sprintf(msg, args...)
		formattedMsg := l.formatMessage(level, msg, fileName, line)
		fmt.Println(formattedMsg)
	}
}

// Save logs to DB
func (l *logger) Save() error {
	fmt.Println("Save logs")
	return nil
}

// isLocalHost returns true if the current host is localhost.
func isLocalHost() bool {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return false
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			if ipnet.IP.String() == "127.0.0.1" {
				return true
			}
		}
	}
	return false
}

// SetLogAll sets the logger to log all messages regardless of the log level.
func (l *logger) SetLogAll(logAll bool) {
	l.logAll = logAll
}

// Printf logs a message using a format string and arguments.
func (l *logger) Printf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	if l.level >= InfoLevel {
		l.log(InfoLevel, msg)
	}
}

func (l *logger) formatMessage(level int, msg string, file string, line int) string {
	color := l.levelColors[level]
	if color == "" {
		color = "\033[0m" // reset
	}
	var fieldsStr string
	for k, v := range l.fields {
		fieldsStr += fmt.Sprintf("[%s=%v]", k, v)
	}
	return fmt.Sprintf("%s[%s]%s %s:%d %s %s\033[0m", color, levelNames[level], color, file, line, fieldsStr, msg)
}

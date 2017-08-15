package logging

import (
	"errors"
	"fmt"
	"io"
	"os"

	logging "github.com/op/go-logging"
)

// LOG 日志记录器对象
var log *logging.Logger

// Password 需要记录密码时可以使用 Password(字符串变量/字符串常量)
type Password string

// Redacted 日志框架会调用这个方法将明文密码转换为其它字符
func (p Password) Redacted() (text interface{}) {
	text = logging.Redact(string(p))
	return
}

// InitLogger 初始化日志框架；需要在启动应用时调用一次
func InitLogger(logFile string, logFormat string, level string) (err error) {
	if nil != log {
		err = errors.New("The log has been initialized")
	}
	backend, err := buildBackend(logFile, buildFormatter(logFormat))
	backend, err = setLevel(backend, level)
	if nil != err {
		return
	}
	logging.SetBackend(backend)
	log = logging.MustGetLogger("HUSKY")
	log.ExtraCalldepth = 1
	return
}

// Error logs a message using ERROR as log level.
func Error(args ...interface{}) {
	log.Error(args...)
}

// Errorf logs a message using ERROR as log level.
func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

// Warning logs a message using WARNING as log level.
func Warning(args ...interface{}) {
	log.Warning(args...)
}

// Warningf logs a message using WARNING as log level.
func Warningf(format string, args ...interface{}) {
	log.Warningf(format, args...)
}

// Notice logs a message using NOTICE as log level.
func Notice(args ...interface{}) {
	log.Notice(args...)
}

// Noticef logs a message using NOTICE as log level.
func Noticef(format string, args ...interface{}) {
	log.Noticef(format, args...)
}

// Info logs a message using INFO as log level.
func Info(args ...interface{}) {
	log.Info(args...)
}

// Infof logs a message using INFO as log level.
func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

// Debug logs a message using DEBUG as log level.
func Debug(args ...interface{}) {
	log.Debug(args...)
}

// Debugf logs a message using DEBUG as log level.
func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

func buildBackend(filePath string, formatter logging.Formatter) (backend logging.Backend, err error) {
	out, err := openFile(filePath)
	if nil != err {
		return
	}
	backend = logging.NewLogBackend(out, "", 0)
	backend = logging.NewBackendFormatter(backend, formatter)
	return
}

func setLevel(_backend logging.Backend, level string) (backend logging.Backend, err error) {
	backend = logging.AddModuleLevel(_backend)
	if backendLeveled, isOk := backend.(logging.LeveledBackend); isOk {
		switch level {
		case LevelCrit:
			backendLeveled.SetLevel(logging.CRITICAL, "")
		case LevelCritical:
			backendLeveled.SetLevel(logging.CRITICAL, "")
		case LevelErro:
			backendLeveled.SetLevel(logging.ERROR, "")
		case LevelError:
			backendLeveled.SetLevel(logging.ERROR, "")
		case LevelWarn:
			backendLeveled.SetLevel(logging.WARNING, "")
		case LevelWarning:
			backendLeveled.SetLevel(logging.WARNING, "")
		case LevelNoti:
			backendLeveled.SetLevel(logging.NOTICE, "")
		case LevelNotice:
			backendLeveled.SetLevel(logging.NOTICE, "")
		case LevelInfo:
			backendLeveled.SetLevel(logging.INFO, "")
		case LevelDebu:
			backendLeveled.SetLevel(logging.DEBUG, "")
		case LevelDebug:
			backendLeveled.SetLevel(logging.DEBUG, "")
		default:
			err = errors.New("日志级别设置错误: " + level)
		}
	}
	return
}

func buildFormatter(logFormat string) (formater logging.Formatter) {
	return logging.MustStringFormatter(logFormat)
}

func openFile(filePath string) (out io.Writer, err error) {
	if "" == filePath {
		out = os.Stdout
		return
	}
	out, err = os.Create(filePath)
	if nil != err {
		fmt.Println("ERROR : Can not open file : " + filePath)
		return
	}
	return
}

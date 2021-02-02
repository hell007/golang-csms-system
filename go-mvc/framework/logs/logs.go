package logs

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"go-mvc/framework/conf"
	"go-mvc/framework/utils/files"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

var (
	once          sync.Once
	logger        *Logger
	fileInfoField string // 文件名和行号 显示字段
)

// 要写入日志的数据字段
type D map[string]interface{}

type Option struct {
	// log 路径
	LogPath string
	// 日志类型 json|text
	LogType string
	// 文件名的日期格式
	FileNameDateFormat string
	// 日志中日期时间格式
	TimestampFormat string
	// 日志级别
	LogLevel string
	// 控制台打印日志
	LogsOut bool
	// 日志最长保存多久
	MaxAge time.Duration
	// 日志默认多长时间轮转一次
	RotationTime time.Duration
	// 是否开启记录文件名和行号
	EnableRecordFileInfo bool
	// 文件名和行号字段名
	FileInfoField string
	// json日志是否美化输出
	JSONPrettyPrint bool
	// json日志条目中 数据字段都会作为该字段的嵌入字段
	JSONDataKey string
}

// 封装logrus结构体
type Logger struct {
	logrus               *logrus.Logger
	enableRecordFileinfo bool
}

// logs日志启动
func Start(p string) {
	once.Do(func() {
		logger, err := New(&Option{
			LogPath:              p,
			LogType:              conf.GlobalConfig.LogsType,
			FileNameDateFormat:   "%Y%m%d",
			TimestampFormat:      conf.GlobalConfig.Timeformat,
			LogsOut:              conf.GlobalConfig.LogsOut,
			LogLevel:             conf.GlobalConfig.LogsLevel,
			MaxAge:               time.Duration(conf.GlobalConfig.LogsMaxAge) * time.Hour,
			RotationTime:         time.Duration(conf.GlobalConfig.LogsRotationTime) * time.Hour,
			JSONPrettyPrint:      conf.GlobalConfig.LogsJSONPrettyPrint,
			JSONDataKey:          conf.GlobalConfig.LogsJSONDataKey,
			EnableRecordFileInfo: conf.GlobalConfig.LogsEnableRecordFileInfo,
			FileInfoField:        conf.GlobalConfig.LogsFileInfoField,
		})

		if err != nil {
			StderrFatalf("error: %s", err)
			return
		}

		logger.Info(nil, "logs日志记录已启动...")
	})
}

//全局获取logger实例
func GetLogger() *Logger {
	return logger
}

func newLogger(option *Option) (*logrus.Logger, error) {
	var (
		err          error
		logrusLogger *logrus.Logger
	)

	if err = files.EnsureDir(option.LogPath); err != nil {
		return nil, err
	}

	fileInfoField = option.FileInfoField

	logrusLogger = logrus.New()

	if option.LogsOut {
		logrusLogger.SetOutput(os.Stdout)
	} else {
		logrusLogger.SetOutput(ioutil.Discard)
	}

	level, _ := logrus.ParseLevel(option.LogLevel)
	logrusLogger.Level = level

	switch option.LogType {
	case "JSON":
		format := &logrus.JSONFormatter{
			TimestampFormat: option.TimestampFormat,
			PrettyPrint:     option.JSONPrettyPrint,
		}
		if option.JSONDataKey != "" {
			format.DataKey = option.JSONDataKey
		}
		logrusLogger.Formatter = format
	default:
		logrusLogger.Formatter = &logrus.TextFormatter{
			TimestampFormat: option.TimestampFormat,
		}
	}

	return logrusLogger, nil
}

func New(option *Option) (*Logger, error) {
	var (
		err          error
		logrusLogger *logrus.Logger
		writer       *rotatelogs.RotateLogs
		fileHook     *lfshook.LfsHook
		absPath      string
	)

	if logrusLogger, err = newLogger(option); err != nil {
		return nil, err
	}

	if isWindow() {
		writer, err = rotatelogs.New(
			fmt.Sprintf("%sjie-%s.log", option.LogPath, option.FileNameDateFormat),
			rotatelogs.WithMaxAge(option.MaxAge),
			rotatelogs.WithRotationTime(option.RotationTime),
		)
	} else {
		absPath, err = filepath.Abs(option.LogPath)
		if err != nil {
			return nil, fmt.Errorf("rotatelogs.New error: %s", err)
		}
		writer, err = rotatelogs.New(
			fmt.Sprintf("%sjie--%s.log", absPath, option.FileNameDateFormat),
			rotatelogs.WithMaxAge(option.MaxAge),
			rotatelogs.WithRotationTime(option.RotationTime),
			rotatelogs.WithLinkName(absPath),
		)
	}

	if err != nil {
		return nil, err
	}

	fileHook = lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, logrusLogger.Formatter)

	logrusLogger.Hooks.Add(fileHook)

	logger = &Logger{
		logrus:               logrusLogger,
		enableRecordFileinfo: option.EnableRecordFileInfo,
	}
	return logger, nil
}

func newRotatelog(option *Option, levelStr string) (*rotatelogs.RotateLogs, error) {
	var (
		err      error
		filename string
		writer   *rotatelogs.RotateLogs
		absPath  string
	)

	filename = fmt.Sprintf("%s.%s", option.LogPath, levelStr)
	if isWindow() {
		writer, err = rotatelogs.New(
			fmt.Sprintf("%s.%s", filename, option.FileNameDateFormat),
			rotatelogs.WithMaxAge(option.MaxAge),
			rotatelogs.WithRotationTime(option.RotationTime),
		)
	} else {
		absPath, err = filepath.Abs(filename)
		if err != nil {
			return nil, fmt.Errorf("rotatelogs.New error: %s", err)
		}

		writer, err = rotatelogs.New(
			fmt.Sprintf("%s.%s", absPath, option.FileNameDateFormat),
			rotatelogs.WithMaxAge(option.MaxAge),
			rotatelogs.WithRotationTime(option.RotationTime),
			rotatelogs.WithLinkName(absPath),
		)
	}

	if err != nil {
		return nil, fmt.Errorf("rotatelogs.New error: %s", err)
	}

	return writer, nil
}

// NewSeparate 不同级别的日志输出到不同的文件
func NewSeparate(option *Option) (*Logger, error) {
	var (
		err          error
		logrusLogger *logrus.Logger
		debugWriter  *rotatelogs.RotateLogs
		infoWriter   *rotatelogs.RotateLogs
		warnWriter   *rotatelogs.RotateLogs
		errorWriter  *rotatelogs.RotateLogs
		fatalWriter  *rotatelogs.RotateLogs
		panicWriter  *rotatelogs.RotateLogs
		fileHook     *lfshook.LfsHook
	)

	if logrusLogger, err = newLogger(option); err != nil {
		return nil, err
	}

	if debugWriter, err = newRotatelog(option, "debug"); err != nil {
		return nil, err
	}

	if infoWriter, err = newRotatelog(option, "info"); err != nil {
		return nil, err
	}

	if warnWriter, err = newRotatelog(option, "warn"); err != nil {
		return nil, err
	}

	if errorWriter, err = newRotatelog(option, "error"); err != nil {
		return nil, err
	}

	if fatalWriter, err = newRotatelog(option, "fatal"); err != nil {
		return nil, err
	}

	if panicWriter, err = newRotatelog(option, "panic"); err != nil {
		return nil, err
	}

	fileHook = lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: debugWriter, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  infoWriter,
		logrus.WarnLevel:  warnWriter,
		logrus.ErrorLevel: errorWriter,
		logrus.FatalLevel: fatalWriter,
		logrus.PanicLevel: panicWriter,
	}, logrusLogger.Formatter)

	logrusLogger.Hooks.Add(fileHook)

	logger = &Logger{
		logrus:               logrusLogger,
		enableRecordFileinfo: option.EnableRecordFileInfo,
	}

	return logger, nil
}

func (l *Logger) Debug(dataFields D, format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	if dataFields == nil {
		dataFields = D{}
	}
	if l.enableRecordFileinfo {
		dataFields[fileInfoField] = fileInfo(2)
	}
	l.logrus.WithFields(logrus.Fields(dataFields)).Debug(message)
}

func (l *Logger) Info(dataFields D, format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	if dataFields == nil {
		dataFields = D{}
	}
	if l.enableRecordFileinfo {
		dataFields[fileInfoField] = fileInfo(2)
	}
	l.logrus.WithFields(logrus.Fields(dataFields)).Info(message)
}

func (l *Logger) Warn(dataFields D, format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	if dataFields == nil {
		dataFields = D{}
	}
	if l.enableRecordFileinfo {
		dataFields[fileInfoField] = fileInfo(2)
	}
	l.logrus.WithFields(logrus.Fields(dataFields)).Warn(message)
}

func (l *Logger) Error(dataFields D, format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	if dataFields == nil {
		dataFields = D{}
	}
	if l.enableRecordFileinfo {
		dataFields[fileInfoField] = fileInfo(2)
	}
	l.logrus.WithFields(logrus.Fields(dataFields)).Error(message)
}

func (l *Logger) Fatal(dataFields D, format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	if dataFields == nil {
		dataFields = D{}
	}
	if l.enableRecordFileinfo {
		dataFields[fileInfoField] = fileInfo(2)
	}
	l.logrus.WithFields(logrus.Fields(dataFields)).Fatal(message)
}

func (l *Logger) Panic(dataFields D, format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	if dataFields == nil {
		dataFields = D{}
	}
	if l.enableRecordFileinfo {
		dataFields[fileInfoField] = fileInfo(2)
	}
	l.logrus.WithFields(logrus.Fields(dataFields)).Panic(message)
}

func Debug(dataFields D, format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	if dataFields == nil {
		dataFields = D{}
	}
	if logger.logrus == nil {
		logrus.Debug(dataFields, message)
		return
	}
	if logger.enableRecordFileinfo {
		dataFields[fileInfoField] = fileInfo(2)
	}
	logger.logrus.WithFields(logrus.Fields(dataFields)).Debug(message)
}

func Info(dataFields D, format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	if dataFields == nil {
		dataFields = D{}
	}
	if logger.logrus == nil {
		logrus.Info(dataFields, message)
		return
	}
	if logger.enableRecordFileinfo {
		dataFields[fileInfoField] = fileInfo(2)
	}
	logger.logrus.WithFields(logrus.Fields(dataFields)).Info(message)
}

func Warn(dataFields D, format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	if dataFields == nil {
		dataFields = D{}
	}
	if logger.logrus == nil {
		logrus.Warn(dataFields, message)
		return
	}
	if logger.enableRecordFileinfo {
		dataFields[fileInfoField] = fileInfo(2)
	}
	logger.logrus.WithFields(logrus.Fields(dataFields)).Warn(message)
}

func Error(dataFields D, format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	if dataFields == nil {
		dataFields = D{}
	}
	if logger.logrus == nil {
		logrus.Error(dataFields, message)
		return
	}
	if logger.enableRecordFileinfo {
		dataFields[fileInfoField] = fileInfo(2)
	}
	logger.logrus.WithFields(logrus.Fields(dataFields)).Error(message)
}

func Fatal(dataFields D, format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	if dataFields == nil {
		dataFields = D{}
	}
	if logger.logrus == nil {
		logrus.Fatal(dataFields, message)
		return
	}
	if logger.enableRecordFileinfo {
		dataFields[fileInfoField] = fileInfo(2)
	}
	logger.logrus.WithFields(logrus.Fields(dataFields)).Fatal(message)
}

func Panic(dataFields D, format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	if dataFields == nil {
		dataFields = D{}
	}
	if logger.logrus == nil {
		logrus.Panic(dataFields, message)
		return
	}
	if logger.enableRecordFileinfo {
		dataFields[fileInfoField] = fileInfo(2)
	}
	logger.logrus.WithFields(logrus.Fields(dataFields)).Panic(message)
}

func StderrFatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}

func fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}

func isWindow() bool {
	return runtime.GOOS == "windows"
}

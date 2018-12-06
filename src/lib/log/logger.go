package log

import (
	"context"
	"fmt"
	"runtime"
	"strings"

	"google.golang.org/appengine/log"
)

// Logger ... ロガー
type Logger struct {
	MinLevel Level
}

// IsLogging ... レベル毎のログ出力許可
func (l *Logger) IsLogging(level Level) bool {
	return l.MinLevel <= level
}

// NewLogger ... Loggerを作成する
func NewLogger(minLevel Level) Logger {
	return Logger{
		MinLevel: minLevel,
	}
}

// Debugf ... Debugログを出力する
func Debugf(ctx context.Context, format string, args ...interface{}) {
	logger := GetLogger(ctx)
	if logger.IsLogging(LevelDebug) {
		fl := getFileLine()
		log.Debugf(ctx, fl+format, args...)
	}
}

// Debugm ... Debugログの定形を出力する
func Debugm(ctx context.Context, method string, err error) {
	logger := GetLogger(ctx)
	if logger.IsLogging(LevelDebug) {
		fl := getFileLine()
		log.Debugf(ctx, "%s%s: %s", fl, method, err.Error())
	}
}

// Infof ... Infoログを出力する
func Infof(ctx context.Context, format string, args ...interface{}) {
	logger := GetLogger(ctx)
	if logger.IsLogging(LevelInfo) {
		fl := getFileLine()
		log.Infof(ctx, fl+format, args...)
	}
}

// Infom ... Infoログの定形を出力する
func Infom(ctx context.Context, method string, err error) {
	logger := GetLogger(ctx)
	if logger.IsLogging(LevelInfo) {
		fl := getFileLine()
		log.Infof(ctx, "%s%s: %s", fl, method, err.Error())
	}
}

// Warningf ... Warningログを出力する
func Warningf(ctx context.Context, format string, args ...interface{}) {
	logger := GetLogger(ctx)
	if logger.IsLogging(LevelWarning) {
		fl := getFileLine()
		log.Warningf(ctx, fl+format, args...)
	}
}

// Warningm ... Warningログの定形を出力する
func Warningm(ctx context.Context, method string, err error) {
	logger := GetLogger(ctx)
	if logger.IsLogging(LevelWarning) {
		fl := getFileLine()
		log.Warningf(ctx, "%s%s: %s", fl, method, err.Error())
	}
}

// Errorf ... Errorログを出力する
func Errorf(ctx context.Context, format string, args ...interface{}) {
	logger := GetLogger(ctx)
	if logger.IsLogging(LevelError) {
		fl := getFileLine()
		log.Errorf(ctx, fl+format, args...)
	}
}

// Errorm ... Errorログの定形を出力する
func Errorm(ctx context.Context, method string, err error) {
	logger := GetLogger(ctx)
	if logger.IsLogging(LevelError) {
		fl := getFileLine()
		log.Errorf(ctx, "%s%s: %s", fl, method, err.Error())
	}
}

// Criticalf ... Criticalログを出力する
func Criticalf(ctx context.Context, format string, args ...interface{}) {
	logger := GetLogger(ctx)
	if logger.IsLogging(LevelCritical) {
		fl := getFileLine()
		log.Criticalf(ctx, fl+format, args...)
	}
}

// Criticalm ... Criticalログの定形を出力する
func Criticalm(ctx context.Context, method string, err error) {
	logger := GetLogger(ctx)
	if logger.IsLogging(LevelCritical) {
		fl := getFileLine()
		log.Criticalf(ctx, "%s%s: %s", fl, method, err.Error())
	}
}

func getFileLine() string {
	var ret string
	if _, file, line, ok := runtime.Caller(2); ok {
		parts := strings.Split(file, "/")
		length := len(parts)
		ret = fmt.Sprintf("%s/%s:%d ", parts[length-2], parts[length-1], line)
	}
	return ret
}

package xlog

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

//todo: the stick has problom

type LoggerLevel int

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
)

var logLevelName = map[LoggerLevel]string{
	DEBUG: "DEBUG",
	INFO:  "INFO",
	WARN:  "WARN",
	ERROR: "ERROR",
}

type Logger struct {
	Log *os.File
}

func NewLogger() (*Logger, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("failed to get current dir")
	}
	filepath := filepath.Dir(filename)

	file, err := os.OpenFile(filepath+"/log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	return &Logger{Log: file}, nil
}

func (l *Logger) Write(level LoggerLevel, args ...any) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	message := fmt.Sprintf("[%s] [%s]", timestamp, logLevelName[level])

	for _, arg := range args {
		message += fmt.Sprintf(" %v", arg)
	}

	fmt.Println(message)

	if l.Log != nil {
		fmt.Fprintln(l.Log, message)
	}
}

func (l *Logger) Close() {
	if l.Log != nil {
		l.Log.Close()
	}
}

func (l *Logger) DEBUG(args ...any) {
	_, file, line, _ := runtime.Caller(1)
	location := fmt.Sprintf("[%s:%d]", file, line)
	args = append(args, location)
	l.Write(DEBUG, args...)
}

func (l *Logger) INFO(args ...any) {
	// _, file, line, _ := runtime.Caller(1)
	// location := fmt.Sprintf("[%s:%d]", file, line)
	// args = append(args)
	l.Write(INFO, args...)
}

func (l *Logger) WARN(args ...any) {
	_, file, line, _ := runtime.Caller(1)
	location := fmt.Sprintf("[%s:%d]", file, line)
	args = append(args, location)
	l.Write(WARN, args...)
}

func (l *Logger) ERROR(args ...any) {
	_, file, line, _ := runtime.Caller(1)
	location := fmt.Sprintf("[%s:%d]", file, line)
	args = append(args, location)
	l.Write(ERROR, args...)
}

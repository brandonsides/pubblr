package logging

import (
	"fmt"
	"io"
	"os"
)

type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warn
	Error
	Fatal
)

type PubblrLogger struct {
	out   io.Writer
	err   io.Writer
	Level LogLevel
}

type PubblrLoggerConfig struct {
	Level LogLevel `json:"level"`
}

func NewStandardPubblrLogger(config PubblrLoggerConfig) *PubblrLogger {
	return NewPubblrLogger(config, os.Stdout, os.Stderr)
}

func NewPubblrLogger(config PubblrLoggerConfig, out io.Writer, err io.Writer) *PubblrLogger {
	return &PubblrLogger{
		out:   out,
		err:   err,
		Level: config.Level,
	}
}

func (l *PubblrLogger) Debugf(format string, args ...interface{}) {
	fmt.Fprintf(l.out, format, args...)
}

func (l *PubblrLogger) Infof(format string, args ...interface{}) {
	fmt.Fprintf(l.out, format, args...)
}

func (l *PubblrLogger) Warnf(format string, args ...interface{}) {
	fmt.Fprintf(l.out, format, args...)
}

func (l *PubblrLogger) Errorf(format string, args ...interface{}) {
	fmt.Fprintf(l.err, format, args...)
}

func (l *PubblrLogger) Fatalf(format string, args ...interface{}) {
	fmt.Fprintf(l.err, format, args...)

}

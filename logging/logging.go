package logging

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warn
	Error
	Fatal
)

func (l LogLevel) String() string {
	switch l {
	case Debug:
		return "DEBUG"
	case Info:
		return "INFO"
	case Warn:
		return "WARN"
	case Error:
		return "ERROR"
	case Fatal:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

func (l *LogLevel) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	err := unmarshal(&str)
	if err != nil {
		return err
	}

	str = strings.ToUpper(str)

	switch str {
	case "DEBUG":
		*l = Debug
	case "INFO":
		*l = Info
	case "WARN":
		*l = Warn
	case "ERROR":
		*l = Error
	case "FATAL":
		*l = Fatal
	default:
		*l = Debug
	}

	return nil
}

type PubblrLogger struct {
	out   io.Writer
	err   io.Writer
	Level LogLevel
}

type PubblrLoggerConfig struct {
	Level LogLevel `yaml:"level"`
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

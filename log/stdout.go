package log

import (
	"fmt"
	"os"

	"bookstore/config"
	"bookstore/time"
)

type stdOutLogger struct {
	time     *time.Factory
	loglevel Loglevel
}

func (s *stdOutLogger) Error(log string, err error) {
	s.logError(ErrorLogLevel, log, err)
}

func (s *stdOutLogger) Debug(log string) {
	s.log(DebugLogLevel, log)
}

func (s *stdOutLogger) Info(log string) {
	s.log(InfoLogLevel, log)
}

func (s *stdOutLogger) log(level Loglevel, log string) {
	fmt.Fprintf(os.Stdout, "\n[%s][%s] %s", logLevelToString(level), s.time.Now().String(), log)
}

func (s *stdOutLogger) logError(level Loglevel, log string, err error) {
	fmt.Fprintf(os.Stdout, "\n[%s][%s] %s error: %s", logLevelToString(level), s.time.Now().String(), log, err)
}

func newStdoutLogger(config *config.LogConfig, time *time.Factory) Backend {
	return &stdOutLogger{
		time:     time,
		loglevel: config.LogLevel,
	}
}

package log

import (
	"fmt"
	"os"

	"bookstore/config"
)

type stdOutLogger struct {
	loglevel Loglevel
}

func (s *stdOutLogger) Debug(log string) {
	s.log(DebugLogLevel, log)
}

func (s *stdOutLogger) Info(log string) {
	s.log(InfoLogLevel, log)
}

func (s *stdOutLogger) log(level Loglevel, log string) {
	fmt.Fprintf(os.Stdout, "[%s] %s", logLevelToString(level), log)
}

func newStdoutLogger(config *config.LogConfig) Backend {
	return &stdOutLogger{
		loglevel: config.LogLevel,
	}
}

package log

type Loglevel = int

const (
	DebugLogLevel Loglevel = iota
	InfoLogLevel
	WarningLoglevel
	ErrorLogLevel
	CriticalLogLevel
)

var logLevelToStringMap = map[Loglevel]string{
	DebugLogLevel:    "DEBUG",
	InfoLogLevel:     "INFO",
	WarningLoglevel:  "WARNING",
	ErrorLogLevel:    "ERROR",
	CriticalLogLevel: "CRITICAL",
}

func logLevelToString(l Loglevel) string {
	return logLevelToStringMap[l]
}

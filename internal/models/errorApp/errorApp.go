package errorapp

import "errors"

//Config
var (
	ErrParseLoggerConfig   = errors.New("failed to parse logger config")
	ErrParsePostgresConfig = errors.New("failed to parse postgres config")
	ErrParseServerConfig   = errors.New("failed to parse server config")
	ErrLoadEnv             = errors.New("failed to load env file")
)

//logger
var (
	ErrFailedOpenLogFile = errors.New("failed to open log file ")
)

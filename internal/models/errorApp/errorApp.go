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

//DataBase
var (
	ErrParseDatabaseConfig = errors.New("failed parse database config")
	ErrCreateNewConnect    = errors.New("failed to create new connect with postgres")
	ErrConnectDatabase     = errors.New("failed to connect with postgres")
)

//Migrator
var (
	ErrSetGooseDialect = errors.New("failed to set goose dialect")
	ErrCloseDb         = errors.New("failed to close db connection")
	ErrUpMigration     = errors.New("failed to up migration")
	ErrUpToMigration   = errors.New("failed to upto migration")
	ErrDownMigration   = errors.New("failed to down migration")
	ErrDownToMigration = errors.New("failed to downto migration")
	ErrCreateMigration = errors.New("failed to create migration")
	ErrGetDbVersion    = errors.New("failed to get db version")
	ErrGetGooseStatus  = errors.New("failed to get goose status")
)

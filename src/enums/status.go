package enums

const (
	StatusNotRunning = "not running"
	StatusRunning    = "running"
	StatusTerminated = "terminated"
)

const (
	GET  = 1
	POST = 2
)

const (
	EndedGracefully      = 0
	GeneralErrorOccurred = 1
	UnhandledException   = 255
)

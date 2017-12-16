package log

import (
	stdLog "log"
)

// Enabled controls logging from crossdock tests. It is enabled in main.go, but off in unit tests.
var Enabled bool

//Debugged enable output debug log
var Debugged bool

// Printf delegates to log.Printf if Enabled == true
func Printf(msg string, args ...interface{}) {
	if Enabled {
		stdLog.Printf(msg, args)
	}
}

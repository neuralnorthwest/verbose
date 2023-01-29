package verbose

import "fmt"

const (
	// Verbose levels
	VerboseLevelNone = iota
	VerboseLevelVerbose
	VerboseLevelDebug
	VerboseLevelTrace
)

// Verbose is an interface for verbose output.
type Verbose interface {
	// Verbosef prints a verbose message.
	Verbosef(format string, args ...interface{})
	// Verbose prints a verbose message.
	Verbose(args ...interface{})
	// Debugf prints a debug message.
	Debugf(format string, args ...interface{})
	// Debug prints a debug message.
	Debug(args ...interface{})
	// Tracef prints a trace message.
	Tracef(format string, args ...interface{})
	// Trace prints a trace message.
	Trace(args ...interface{})
}

// verbose is the verbose implementation.
type verbose struct {
	// VerboseLevel is the verbose level.
	VerboseLevel int
}

// New creates a new verbose implementation.
func New(verboseLevel int) Verbose {
	return &verbose{
		VerboseLevel: verboseLevel,
	}
}

// Verbosef prints a verbose message.
func (v *verbose) Verbosef(format string, args ...interface{}) {
	if v.VerboseLevel >= VerboseLevelVerbose {
		fmt.Printf(format, args...)
	}
}

// Verbose prints a verbose message.
func (v *verbose) Verbose(args ...interface{}) {
	if v.VerboseLevel >= VerboseLevelVerbose {
		fmt.Println(args...)
	}
}

// Debugf prints a debug message.
func (v *verbose) Debugf(format string, args ...interface{}) {
	if v.VerboseLevel >= VerboseLevelDebug {
		fmt.Printf(format, args...)
	}
}

// Debug prints a debug message.
func (v *verbose) Debug(args ...interface{}) {
	if v.VerboseLevel >= VerboseLevelDebug {
		fmt.Println(args...)
	}
}

// Tracef prints a trace message.
func (v *verbose) Tracef(format string, args ...interface{}) {
	if v.VerboseLevel >= VerboseLevelTrace {
		fmt.Printf(format, args...)
	}
}

// Trace prints a trace message.
func (v *verbose) Trace(args ...interface{}) {
	if v.VerboseLevel >= VerboseLevelTrace {
		fmt.Println(args...)
	}
}

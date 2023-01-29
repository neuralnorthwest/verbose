package verbose

import (
	"fmt"
	"io"
	"os"
)

const (
	// Verbose levels
	LevelNone = iota
	LevelVerbose
	LevelDebug
	LevelTrace
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

// VerboseImpl is the VerboseImpl implementation.
type VerboseImpl struct {
	// verboseLevel is the verbose level.
	verboseLevel int
	// writer is the writer.
	writer io.Writer
	// printLevelPrefix is a flag that indicates whether the verbose level
	// prefix should be printed.
	printLevelPrefix bool
}

// VerboseOption is an option for verbose.
type VerboseOption func(*VerboseImpl)

// WithWriter sets the writer.
func WithWriter(w io.Writer) VerboseOption {
	return func(v *VerboseImpl) {
		v.writer = w
	}
}

// WithLevelPrefix is a VerboseOption that prints the verbose level prefix.
func WithLevelPrefix() VerboseOption {
	return func(v *VerboseImpl) {
		v.printLevelPrefix = true
	}
}

// New creates a new verbose implementation.
func New(verboseLevel int, opts ...VerboseOption) Verbose {
	// Limit the verbose level to the valid range.
	if verboseLevel < LevelNone {
		verboseLevel = LevelNone
	} else if verboseLevel > LevelTrace {
		verboseLevel = LevelTrace
	}
	v := &VerboseImpl{
		verboseLevel: verboseLevel,
		writer:       os.Stdout,
	}
	for _, opt := range opts {
		opt(v)
	}
	return v
}

// Verbosef prints a verbose message. Verbose messages are only printed if the
// verbose level is set to LevelVerbose or higher.
func (v *VerboseImpl) Verbosef(format string, args ...interface{}) {
	if v.verboseLevel >= LevelVerbose {
		if v.printLevelPrefix {
			format = "[VERBOSE] " + format
		}
		fmt.Fprintf(v.writer, format, args...)
	}
}

// Verbose prints a verbose message. Verbose messages are only printed if the
// verbose level is set to LevelVerbose or higher.
func (v *VerboseImpl) Verbose(args ...interface{}) {
	if v.verboseLevel >= LevelVerbose {
		if v.printLevelPrefix {
			args = append([]interface{}{"[VERBOSE]"}, args...)
		}
		fmt.Fprintln(v.writer, args...)
	}
}

// Debugf prints a debug message. Debug messages are only printed if the
// verbose level is set to LevelDebug or higher.
func (v *VerboseImpl) Debugf(format string, args ...interface{}) {
	if v.verboseLevel >= LevelDebug {
		if v.printLevelPrefix {
			format = "[DEBUG] " + format
		}
		fmt.Fprintf(v.writer, format, args...)
	}
}

// Debug prints a debug message. Debug messages are only printed if the
// verbose level is set to LevelDebug or higher.
func (v *VerboseImpl) Debug(args ...interface{}) {
	if v.verboseLevel >= LevelDebug {
		if v.printLevelPrefix {
			args = append([]interface{}{"[DEBUG]"}, args...)
		}
		fmt.Fprintln(v.writer, args...)
	}
}

// Tracef prints a trace message. Trace messages are only printed if the
// verbose level is set to LevelTrace.
func (v *VerboseImpl) Tracef(format string, args ...interface{}) {
	if v.verboseLevel >= LevelTrace {
		if v.printLevelPrefix {
			format = "[TRACE] " + format
		}
		fmt.Fprintf(v.writer, format, args...)
	}
}

// Trace prints a trace message. Trace messages are only printed if the
// verbose level is set to LevelTrace.
func (v *VerboseImpl) Trace(args ...interface{}) {
	if v.verboseLevel >= LevelTrace {
		if v.printLevelPrefix {
			args = append([]interface{}{"[TRACE]"}, args...)
		}
		fmt.Fprintln(v.writer, args...)
	}
}

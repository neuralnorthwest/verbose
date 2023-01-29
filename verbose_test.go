package verbose

import (
	"bytes"
	"os"
	"testing"
)

// Test_Verbose_Case is a test case for Verbose.
type Test_Verbose_Case struct {
	// Name is the test name.
	Name string
	// Options are the verbose options.
	Options []VerboseOption
	// VerboseLevel is the verbose level.
	VerboseLevel int
	// Calls are the calls to make.
	Calls []func(Verbose)
	// Expected is the expected output.
	Expected string
}

// Test_Verbose tests Verbose.
func Test_Verbose(t *testing.T) {
	cases := []Test_Verbose_Case{
		{
			Name: "none",
			Options: []VerboseOption{
				WithWriter(&bytes.Buffer{}),
			},
			VerboseLevel: LevelNone,
			Calls: []func(Verbose){
				func(v Verbose) {
					v.Verbosef("verbosef")
				},
				func(v Verbose) {
					v.Verbose("verbose")
				},
				func(v Verbose) {
					v.Debugf("debugf")
				},
				func(v Verbose) {
					v.Debug("debug")
				},
				func(v Verbose) {
					v.Tracef("tracef")
				},
				func(v Verbose) {
					v.Trace("trace")
				},
			},
			Expected: "",
		},
		{
			Name: "verbose",
			Options: []VerboseOption{
				WithWriter(&bytes.Buffer{}),
			},
			VerboseLevel: LevelVerbose,
			Calls: []func(Verbose){
				func(v Verbose) {
					v.Verbosef("verbosef")
				},
				func(v Verbose) {
					v.Verbose("verbose")
				},
				func(v Verbose) {
					v.Debugf("debugf")
				},
				func(v Verbose) {
					v.Debug("debug")
				},
				func(v Verbose) {
					v.Tracef("tracef")
				},
				func(v Verbose) {
					v.Trace("trace")
				},
			},
			Expected: "verbosefverbose\n",
		},
		{
			Name: "debug",
			Options: []VerboseOption{
				WithWriter(&bytes.Buffer{}),
			},
			VerboseLevel: LevelDebug,
			Calls: []func(Verbose){
				func(v Verbose) {
					v.Verbosef("verbosef")
				},
				func(v Verbose) {
					v.Verbose("verbose")
				},
				func(v Verbose) {
					v.Debugf("debugf")
				},
				func(v Verbose) {
					v.Debug("debug")
				},
				func(v Verbose) {
					v.Tracef("tracef")
				},
				func(v Verbose) {
					v.Trace("trace")
				},
			},
			Expected: "verbosefverbose\ndebugfdebug\n",
		},
		{
			Name: "trace",
			Options: []VerboseOption{
				WithWriter(&bytes.Buffer{}),
			},
			VerboseLevel: LevelTrace,
			Calls: []func(Verbose){
				func(v Verbose) {
					v.Verbosef("verbosef")
				},
				func(v Verbose) {
					v.Verbose("verbose")
				},
				func(v Verbose) {
					v.Debugf("debugf")
				},
				func(v Verbose) {
					v.Debug("debug")
				},
				func(v Verbose) {
					v.Tracef("tracef")
				},
				func(v Verbose) {
					v.Trace("trace")
				},
			},
			Expected: "verbosefverbose\ndebugfdebug\ntraceftrace\n",
		},
		{
			Name: "verbose with prefix",
			Options: []VerboseOption{
				WithWriter(&bytes.Buffer{}),
				WithLevelPrefix(),
			},
			VerboseLevel: LevelVerbose,
			Calls: []func(Verbose){
				func(v Verbose) {
					v.Verbosef("verbosef")
				},
				func(v Verbose) {
					v.Verbose("verbose")
				},
				func(v Verbose) {
					v.Debugf("debugf")
				},
				func(v Verbose) {
					v.Debug("debug")
				},
				func(v Verbose) {
					v.Tracef("tracef")
				},
				func(v Verbose) {
					v.Trace("trace")
				},
			},
			Expected: "[VERBOSE] verbosef[VERBOSE] verbose\n",
		},
		{
			Name: "debug with prefix",
			Options: []VerboseOption{
				WithWriter(&bytes.Buffer{}),
				WithLevelPrefix(),
			},
			VerboseLevel: LevelDebug,
			Calls: []func(Verbose){
				func(v Verbose) {
					v.Verbosef("verbosef")
				},
				func(v Verbose) {
					v.Verbose("verbose")
				},
				func(v Verbose) {
					v.Debugf("debugf")
				},
				func(v Verbose) {
					v.Debug("debug")
				},
				func(v Verbose) {
					v.Tracef("tracef")
				},
				func(v Verbose) {
					v.Trace("trace")
				},
			},
			Expected: "[VERBOSE] verbosef[VERBOSE] verbose\n[DEBUG] debugf[DEBUG] debug\n",
		},
		{
			Name: "trace with prefix",
			Options: []VerboseOption{
				WithWriter(&bytes.Buffer{}),
				WithLevelPrefix(),
			},
			VerboseLevel: LevelTrace,
			Calls: []func(Verbose){
				func(v Verbose) {
					v.Verbosef("verbosef")
				},
				func(v Verbose) {
					v.Verbose("verbose")
				},
				func(v Verbose) {
					v.Debugf("debugf")
				},
				func(v Verbose) {
					v.Debug("debug")
				},
				func(v Verbose) {
					v.Tracef("tracef")
				},
				func(v Verbose) {
					v.Trace("trace")
				},
			},
			Expected: "[VERBOSE] verbosef[VERBOSE] verbose\n[DEBUG] debugf[DEBUG] debug\n[TRACE] tracef[TRACE] trace\n",
		},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			v := New(c.VerboseLevel, c.Options...)
			for _, call := range c.Calls {
				call(v)
			}
			actual := v.(*VerboseImpl).writer.(*bytes.Buffer).String()
			if actual != c.Expected {
				t.Errorf("expected %q, got %q", c.Expected, actual)
			}
		})
	}
}

// Test_Verbose_DefaultWriter tests that the default writer is os.Stdout.
func Test_Verbose_DefaultWriter(t *testing.T) {
	v, ok := New(LevelVerbose).(*VerboseImpl)
	if !ok {
		t.Fatal("expected verbose to be of type *verbose")
	}
	if v.writer != os.Stdout {
		t.Errorf("expected writer to be %v, got %v", os.Stdout, v.writer)
	}
}

// Test_VerboseLevel_Bounds tests that the verbose level is bounded to the
// minimum and maximum values.
func Test_VerboseLevel_Bounds(t *testing.T) {
	vMin, ok := New(-1).(*VerboseImpl)
	if !ok {
		t.Fatal("expected verbose to be of type *verbose")
	}
	if vMin.verboseLevel != LevelNone {
		t.Errorf("expected verbose level to be %d, got %d", LevelNone, vMin.verboseLevel)
	}
	vMax, ok := New(4).(*VerboseImpl)
	if !ok {
		t.Fatal("expected verbose to be of type *verbose")
	}
	if vMax.verboseLevel != LevelTrace {
		t.Errorf("expected verbose level to be %d, got %d", LevelTrace, vMax.verboseLevel)
	}
}

// Test_Verbose_Embed tests that embedding VerboseImpl works.
func Test_Verbose_Embed(t *testing.T) {
	type embed struct {
		Verbose
	}
	v := embed{
		Verbose: New(LevelVerbose, WithWriter(&bytes.Buffer{})),
	}
	v.Verbosef("test")
	if v.Verbose.(*VerboseImpl).writer.(*bytes.Buffer).String() != "test" {
		t.Error("expected embedded verbose to work")
	}
}

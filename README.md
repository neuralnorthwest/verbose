# package verbose

This package provides an interface `Verbose` which exports methods for printing
messages at different levels of verbosity.

## Usage

Import the package:

```go
import "github.com/neuralnorthwest/verbose"
```

Create a new `Verbose` instance at the desired verbose level:

```go
v := verbose.New(verbose.LevelVerbose) // or LevelNone, LevelDebug, or LevelTrace
```

Print messages at different levels of verbosity:

```go
v.Verbose("This is a verbose message")
v.Debug("This is a debug message")
v.Trace("This is a trace message")
```

Formatting is supported:

```go
v.Verbosef("This is a %s message", "verbose")
v.Debugf("This is a %s message", "debug")
v.Tracef("This is a %s message", "trace")
```

## Verbose Levels

| Level | Constant       | Description                                                                        |
|-------|----------------|------------------------------------------------------------------------------------|
| 0     | `LevelNone`    | No output. At this level, all messages are ignored.                                |
| 1     | `LevelVerbose` | Verbose output. At this level, only `Verbose` messages are printed.                |
| 2     | `LevelDebug`   | Debug output. At this level, `Verbose` and `Debug` messages are printed.           |
| 3     | `LevelTrace`   | Trace output. At this level, `Verbose`, `Debug`, and `Trace` messages are printed. |

Typically, the verbose level is set by a command line flag. For example:

```go
var verboseLevel = flag.Int("verbose", 0, "verbose level")
```

Then, the verbose level can be passed to `New`:

```go
v := verbose.New(*verboseLevel)
```

## Options

You can pass options to `New` to customize the output.

| Option                          | Description                                                |
|---------------------------------|------------------------------------------------------------|
| `verbose.WithLevelPrefix()`     | Print the level prefix before each message.                |
| `verbose.WithWriter(io.Writer)` | Set the writer to use for output. Defaults to `os.Stdout`. |

### WithWriter

The `WithWriter` option sets the writer to use for output. By default, the
writer is `os.Stdout`.

For example:

```go
v := verbose.New(1, verbose.WithWriter(os.Stderr))
```

### WithLevelPrefix

The `WithLevelPrefix` option prints the level prefix before each message. It
is formatted as `[<level>] `.

For example:

```go
v := verbose.New(verbose.LevelTrace, verbose.WithLevelPrefix())
v.Verbose("This is a verbose message")
v.Debug("This is a debug message")
v.Trace("This is a trace message")
```

Output:

```
[VERBOSE] This is a verbose message
[DEBUG] This is a debug message
[TRACE] This is a trace message
```

## Embedding Verbose into your own types

To ease the use of `Verbose` in your own types, you can embed the `Verbose`
type into your own type. This allows you to use the `Verbose` methods directly
on your type.

For example:

```go
type MyType struct {
    // Embed the Verbose type into your own type.
    verbose.Verbose
}

func NewMyType(verboseLevel int, opts ...verbose.Option) *MyType {
    return &MyType{
        // Initialize the embedded Verbose type.
        Verbose: verbose.New(verboseLevel, opts...),
    }
}

func (m *MyType) DoSomething() {
    // Use the Verbose methods directly on your type.
    m.Verbose("This is a verbose message")
    m.Debug("This is a debug message")
    m.Trace("This is a trace message")
}
```

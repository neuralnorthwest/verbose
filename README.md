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
verboseLevel := LevelVerbose // or LevelDebug or LevelTrace
v := verbose.New(verboseLevel)
```

Print messages at different levels of verbosity:

```go
v.Verbose("This is a verbose message")
v.Debug("This is a debug message")
v.Trace("This is a trace message")
```

## Verbose Levels

| Level | Constant       | Description    |
|-------|----------------|----------------|
| 0     | `LevelNone`    | No output      |
| 1     | `LevelVerbose` | Verbose output |
| 2     | `LevelDebug`   | Debug output   |
| 3     | `LevelTrace`   | Trace output   |

## Options

You can pass options to `New` to customize the output.

| Option                          | Description                                                |
|---------------------------------|------------------------------------------------------------|
| `verbose.PrintLevelPrefix()`    | Print the level prefix before each message.                |
| `verbose.WithWriter(io.Writer)` | Set the writer to use for output. Defaults to `os.Stdout`. |

For example:

```go
v := verbose.New(1, verbose.WithWriter(os.Stderr), verbose.PrintLevelPrefix())
```

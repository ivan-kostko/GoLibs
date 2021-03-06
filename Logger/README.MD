# Logger
--
#### Description

    Package Logger provides a common interface and wrapper implementation for logging libraries.
    Contains predefined loggers: StderrLogger(prints error log to Std out)

NB:

    For the moment the package is in development(see TODOs).
    Currently, due to possible changes, it is recommended not to use anything from the package,
    except ILogger interface.

#### Install

    go get "github.com/ivan-kostko/GoLibs/Logger"

#### Import

    "github.com/ivan-kostko/GoLibs/Logger"

### Usage

##### type ILogger

```go
type ILogger interface {
	// Emergency logs with an emergency level.
	Emergency(args ...interface{})

	// Emergencyf logs with an emergency level.
	// Arguments are handled in the manner of fmt.Printf.
	Emergencyf(format string, args ...interface{})

	// Alert logs with an alert level.
	Alert(args ...interface{})

	// Alertf logs with an alert level.
	// Arguments are handled in the manner of fmt.Printf.
	Alertf(format string, args ...interface{})

	// Critical logs with a critical level.
	Critical(args ...interface{})

	// Criticalf logs with a critical level.
	// Arguments are handled in the manner of fmt.Printf.
	Criticalf(format string, args ...interface{})

	// Error logs with an error level.
	Error(args ...interface{})

	// Errorf logs with an error level.
	// Arguments are handled in the manner of fmt.Printf.
	Errorf(format string, args ...interface{})

	// Warning logs with a warning level.
	Warning(args ...interface{})

	// Warningf logs with a warning level.
	// Arguments are handled in the manner of fmt.Printf.
	Warningf(format string, args ...interface{})

	// Notice logs with a notice level.
	Notice(args ...interface{})

	// Noticef logs with a notice level.
	// Arguments are handled in the manner of fmt.Printf.
	Noticef(format string, args ...interface{})

	// Info logs with an info level.
	Info(args ...interface{})

	// Infof logs with an info level.
	// Arguments are handled in the manner of fmt.Printf.
	Infof(format string, args ...interface{})

	// Debug logs with a debug level.
	Debug(args ...interface{})

	// Debugf logs with a debug level.
	// Arguments are handled in the manner of fmt.Printf.
	Debugf(format string, args ...interface{})

	// Log logs at the level passed in argument.
	Log(level Level, args ...interface{})

	// Logf logs at the level passed in argument.
	// Arguments are handled in the manner of fmt.Printf.
	Logf(level Level, format string, args ...interface{})
}
```

ILogger is a common interface for logging.

##### func  GetILogger

```go
func GetILogger(conf LoggerConfig) ILogger
```
ILogger factory. Instantiates a new LogAdapter based on provided configuration
and returns it as ILogger

NB : For the moment it returns only StdTerminalLogger independent on
configuration.

TODO(me): Refactor GetILogger, LoggerConfig to support some other logging libs.

##### func  GetStdTerminalLogger

```go
func GetStdTerminalLogger() ILogger
```

##### type Level

```go
type Level int8
```

Level specifies a level of severity. The available levels are the eight
severities described in RFC 5424 and None

```go
const (
	None      Level = iota - 1
	Emergency       //      Emergency: system is unusable
	Alert           //      Alert: action must be taken immediately
	Critical        //      Critical: critical conditions
	Error           //      Error: error conditions
	Warning         //      Warning: warning conditions
	Notice          //      Notice: normal but significant condition
	Info            //      Informational: informational messages
	Debug           //      Debug: debug-level messages
)
```

##### func (Level) String

```go
func (i Level) String() string
```

##### type LogAdapter

```go
type LogAdapter struct {
}
```

LogAdapter adapts logging function func(level Level, args ...interface{}) to
ILogger interface. Could be used for mocking and quick simple introduction of
any logger

NB: For production loggers it is better to create its own adapter

##### func  GetNewLogAdapter

```go
func GetNewLogAdapter(intLog func(level Level, args ...interface{})) *LogAdapter
```
LogAdapter factory. Instantiates a new instance of LogAdapter adapting intLog to
ILogger interface

##### func (*LogAdapter) Alert

```go
func (lc *LogAdapter) Alert(args ...interface{})
```
Alert logs with an emergency level

##### func (*LogAdapter) Alertf

```go
func (lc *LogAdapter) Alertf(format string, args ...interface{})
```
Alertf logs with an emergency level. Arguments are handled in the manner of
fmt.Printf.

##### func (*LogAdapter) Critical

```go
func (lc *LogAdapter) Critical(args ...interface{})
```
Critical logs with an emergency level

##### func (*LogAdapter) Criticalf

```go
func (lc *LogAdapter) Criticalf(format string, args ...interface{})
```
Criticalf logs with an emergency level. Arguments are handled in the manner of
fmt.Printf.

##### func (*LogAdapter) Debug

```go
func (lc *LogAdapter) Debug(args ...interface{})
```
Debug logs with an emergency level

##### func (*LogAdapter) Debugf

```go
func (lc *LogAdapter) Debugf(format string, args ...interface{})
```
Debugf logs with an emergency level. Arguments are handled in the manner of
fmt.Printf.

##### func (*LogAdapter) Emergency

```go
func (lc *LogAdapter) Emergency(args ...interface{})
```
Emergency logs with an emergency level

##### func (*LogAdapter) Emergencyf

```go
func (lc *LogAdapter) Emergencyf(format string, args ...interface{})
```
Emergencyf logs with an emergency level. Arguments are handled in the manner of
fmt.Printf.

##### func (*LogAdapter) Error

```go
func (lc *LogAdapter) Error(args ...interface{})
```
Error logs with an emergency level

##### func (*LogAdapter) Errorf

```go
func (lc *LogAdapter) Errorf(format string, args ...interface{})
```
Errorf logs with an emergency level. Arguments are handled in the manner of
fmt.Printf.

##### func (*LogAdapter) Info

```go
func (lc *LogAdapter) Info(args ...interface{})
```
Info logs with an emergency level

##### func (*LogAdapter) Infof

```go
func (lc *LogAdapter) Infof(format string, args ...interface{})
```
Infof logs with an emergency level. Arguments are handled in the manner of
fmt.Printf.

##### func (*LogAdapter) Log

```go
func (lc *LogAdapter) Log(level Level, args ...interface{})
```
Log logs with an emergency level

##### func (*LogAdapter) Logf

```go
func (lc *LogAdapter) Logf(level Level, format string, args ...interface{})
```
Logf logs with an emergency level. Arguments are handled in the manner of
fmt.Printf.

##### func (*LogAdapter) Notice

```go
func (lc *LogAdapter) Notice(args ...interface{})
```
Notice logs with an emergency level

##### func (*LogAdapter) Noticef

```go
func (lc *LogAdapter) Noticef(format string, args ...interface{})
```
Noticef logs with an emergency level. Arguments are handled in the manner of
fmt.Printf.

##### func (*LogAdapter) Warning

```go
func (lc *LogAdapter) Warning(args ...interface{})
```
Warning logs with an emergency level

##### func (*LogAdapter) Warningf

```go
func (lc *LogAdapter) Warningf(format string, args ...interface{})
```
Warningf logs with an emergency level. Arguments are handled in the manner of
fmt.Printf.

##### type LoggerConfig

```go
type LoggerConfig struct {
	Prefix string
}
```

Represents configuration to create a new Logger

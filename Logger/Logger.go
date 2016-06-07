//   Copyright (c) 2016 Ivan A Kostko (github.com/ivan-kostko)

//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at

//       http://www.apache.org/licenses/LICENSE-2.0

//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package Logger

import (
	"fmt"
)

//go:generate stringer -type=Level

// Level specifies a level of severity. The available levels are the eight severities described in RFC 5424 and None
type Level int8

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

// ILogger is a common interface for logging.
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

// LogAdapter adapts logging function func(level Level, args ...interface{}) to ILogger interface.
// Could be used for mocking and quick simple introduction of any logger
//
// NB: For production loggers it is better to create its own adapter
type LogAdapter struct {
	internalLogFunction func(level Level, args ...interface{})
}

// LogAdapter factory.
// Instantiates a new instance of LogAdapter adapting intLog to ILogger interface
func GetNewLogAdapter(intLog func(level Level, args ...interface{})) *LogAdapter {
	lc := new(LogAdapter)
	lc.internalLogFunction = intLog
	return lc
}

// Emergency logs with an emergency level
func (lc *LogAdapter) Emergency(args ...interface{}) {
	lc.Log(Emergency, args...)
}

// Emergencyf logs with an emergency level.
// Arguments are handled in the manner of fmt.Printf.
func (lc *LogAdapter) Emergencyf(format string, args ...interface{}) {
	lc.Log(Emergency, fmt.Sprintf(format, args...))
}

// Alert logs with an emergency level
func (lc *LogAdapter) Alert(args ...interface{}) {
	lc.Log(Alert, args...)
}

// Alertf logs with an emergency level.
// Arguments are handled in the manner of fmt.Printf.
func (lc *LogAdapter) Alertf(format string, args ...interface{}) {
	lc.Log(Alert, fmt.Sprintf(format, args...))
}

// Critical logs with an emergency level
func (lc *LogAdapter) Critical(args ...interface{}) {
	lc.Log(Critical, args...)
}

// Criticalf logs with an emergency level.
// Arguments are handled in the manner of fmt.Printf.
func (lc *LogAdapter) Criticalf(format string, args ...interface{}) {
	lc.Log(Critical, fmt.Sprintf(format, args...))
}

// Error logs with an emergency level
func (lc *LogAdapter) Error(args ...interface{}) {
	lc.Log(Error, args...)
}

// Errorf logs with an emergency level.
// Arguments are handled in the manner of fmt.Printf.
func (lc *LogAdapter) Errorf(format string, args ...interface{}) {
	lc.Log(Error, fmt.Sprintf(format, args...))
}

// Warning logs with an emergency level
func (lc *LogAdapter) Warning(args ...interface{}) {
	lc.Log(Warning, args...)
}

// Warningf logs with an emergency level.
// Arguments are handled in the manner of fmt.Printf.
func (lc *LogAdapter) Warningf(format string, args ...interface{}) {
	lc.Log(Warning, fmt.Sprintf(format, args...))
}

// Notice logs with an emergency level
func (lc *LogAdapter) Notice(args ...interface{}) {
	lc.Log(Notice, args...)
}

// Noticef logs with an emergency level.
// Arguments are handled in the manner of fmt.Printf.
func (lc *LogAdapter) Noticef(format string, args ...interface{}) {
	lc.Log(Notice, fmt.Sprintf(format, args...))
}

// Info logs with an emergency level
func (lc *LogAdapter) Info(args ...interface{}) {
	lc.Log(Info, args...)
}

// Infof logs with an emergency level.
// Arguments are handled in the manner of fmt.Printf.
func (lc *LogAdapter) Infof(format string, args ...interface{}) {
	lc.Log(Info, fmt.Sprintf(format, args...))
}

// Debug logs with an emergency level
func (lc *LogAdapter) Debug(args ...interface{}) {
	lc.Log(Debug, args...)
}

// Debugf logs with an emergency level.
// Arguments are handled in the manner of fmt.Printf.
func (lc *LogAdapter) Debugf(format string, args ...interface{}) {
	lc.Log(Debug, fmt.Sprintf(format, args...))
}

// Log logs with an emergency level
func (lc *LogAdapter) Log(level Level, args ...interface{}) {
	lc.internalLogFunction(level, args...)
}

// Logf logs with an emergency level.
// Arguments are handled in the manner of fmt.Printf.
func (lc *LogAdapter) Logf(level Level, format string, args ...interface{}) {
	lc.Log(level, fmt.Sprintf(format, args...))
}

// Represents configuration to create a new Logger
type LoggerConfig struct {
	Prefix string
}

// ILogger factory.
// Instantiates a new LogAdapter based on provided configuration and returns it as ILogger
//
// NB : For the moment it returns only StdTerminalLogger independent on configuration.
//
// TODO(me): Refactor GetILogger, LoggerConfig to support some other logging libs.
func GetILogger(conf LoggerConfig) ILogger {
	if conf == (LoggerConfig{}) {
		return GetStdTerminalLogger()
	}
	return GetStdTerminalLogger()
}

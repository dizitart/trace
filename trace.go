// Copyright 2014 Dizitart. All rights reserved.
// Use of this source code is governed by a Apache License V2
// that can be found in the LICENSE file.
//
// Author - Anindya Chatterjee (axchatt@dizitart.com)


package trace

import (
	"fmt"
	"time"
	"runtime"
	"os"
	"path/filepath"
	"runtime/debug"
)

// Use these configuration variables to set tracing parameters
var (
	// Flag to enable or disable the Tracing
	ENABLE_TRACE		bool 	= false
	// Flag to set the tracer to print in Stdout
	PRINT_TO_STDOUT		bool 	= true
	// Flag to set the tracer to print in Stdout
	PRINT_TO_FILE		bool 	= false
	// Trace log file path
	TRACE_OUT_FILE		string
	// Trace message format
	TRACE_FORMAT		string	= "TRACE [%v] <%s> (%s:%d) - %v"
)

// Writes information about current source file, line number and input variable
func Write(variable interface {}) {
	if ENABLE_TRACE {
		msg := createTraceMessage(variable)
		trace(msg, false)
	}
}

// Writes information about current source file, line number and input variable after
// formatting it according to a format specifier
func Writef(format string, variables ...interface{}) {
	if ENABLE_TRACE {
		msg := createTraceMessagef(format, variables...)
		trace(msg, false)
	}
}

// Checks for a condition; if the condition is false, writes
// a trace message that shows the call stack.
func Assert (condition bool) {
	if ENABLE_TRACE {
		if !condition {
			msg := createTraceMessage("Assertion Failed!")
			trace(msg, true)
		}
	}
}

// Checks for a condition; if the condition is false, writes a
// specified message after formatting it according to a format specifier
// that shows the call stack.
func Assertf (condition bool, format string, variables ...interface{}) {
	if ENABLE_TRACE {
		if !condition {
			msg := createTraceMessagef(format, variables...)
			trace(msg, true)
		}
	}
}

// Prints the message and stack trace if set to true
func trace(msg string, printStack bool) {
	if PRINT_TO_STDOUT {
		fmt.Println(msg)
		if printStack {
			debug.PrintStack()
		}
	}

	if PRINT_TO_FILE {
		if TRACE_OUT_FILE == "" {
			return
		} else {
			file, _ := os.OpenFile(TRACE_OUT_FILE, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
			defer file.Close()
			file.WriteString(msg + "\n")
			if printStack {
				file.WriteString(string(debug.Stack()))
			}
		}
	}
}

// Creates the trace message
func createTraceMessage(variable interface {}) string {
	pc, path, line, ok := runtime.Caller(2)
	if ok {
		fun := runtime.FuncForPC(pc)
		name := fun.Name()

		sourceFile := filepath.Base(path)
		msg := fmt.Sprintf(TRACE_FORMAT, time.Now(), name, sourceFile, line, variable)
		return msg
	}
	return ""
}

// Creates the trace message after
// formatting it according to a format specifier
func createTraceMessagef(format string, variables ...interface{}) string {
	pc, path, line, ok := runtime.Caller(2)
	if ok {
		fun := runtime.FuncForPC(pc)
		name := fun.Name()

		sourceFile := filepath.Base(path)
		umsg := fmt.Sprintf(format, variables...)
		msg := fmt.Sprintf(TRACE_FORMAT, time.Now(), name, sourceFile, line, umsg)
		return msg
	}
	return ""
}

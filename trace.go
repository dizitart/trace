// Copyright 2014 Dizitart. All rights reserved.
// Use of this source code is governed by a Apache License V2
// that can be found in the LICENSE file.
//
// Author - Anindya Chatterjee (anindya@dizitart.com)


package trace

import (
	"fmt"
	"time"
	"runtime"
	"os"
	"path/filepath"
	"runtime/debug"
)

type MessageFormat	byte
type OutputChannel	byte

const (
	// Switches off other information in trace message
	PRINT_NONE		MessageFormat = 0
	// Switch to print time
	PRINT_TIME		MessageFormat = 1
	// Switch to print function name
	PRINT_PROC		MessageFormat = 2
	// Switch to print source file name
	PRINT_FILE		MessageFormat = 4
	// Switch yo print current line number
	PRINT_LINE		MessageFormat = 8
	// Switches on all formatting
	PRINT_ALL		MessageFormat = PRINT_TIME | PRINT_PROC | PRINT_FILE | PRINT_LINE

	// Switches off printing
	OUT_NONE		OutputChannel = 0
	// Switch to print to Stdout
	OUT_STD			OutputChannel = 1
	// Switch to print to File
	OUT_FILE		OutputChannel = 2
	// Switches on both File & Stdout printing
	OUT_ALL			OutputChannel = OUT_STD | OUT_FILE
)

// Use these configuration variables to set tracing parameters
var (
	// Flag to enable or disable the Tracing
	ENABLE_TRACE		bool 	= false
	// Trace log file path
	TRACE_FILE_PATH		string
	// Trace message flag
	traceMessageFlag	MessageFormat
	// Output flag
	traceOutFlag		OutputChannel
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

// Sets output channels (Stdout, File)
func SetOut(out OutputChannel) {
	traceOutFlag = out
}

// Sets trace message format
func SetMessageFormat(messageFormat MessageFormat) {
	traceMessageFlag = messageFormat
}

// Prints the message and stack trace if set to true
func trace(msg string, printStack bool) {
	if traceOutFlag & OUT_STD == OUT_STD {
		fmt.Println(msg)
		if printStack {
			debug.PrintStack()
		}
	}

	if traceOutFlag & OUT_FILE == OUT_FILE {
		if TRACE_FILE_PATH == "" {
			return
		} else {
			file, _ := os.OpenFile(TRACE_FILE_PATH, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
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

		var msg string = "TRACE"
		if traceMessageFlag & PRINT_NONE == PRINT_NONE {
			msg = msg + ""
		}
		if traceMessageFlag & PRINT_TIME == PRINT_TIME {
			msg = fmt.Sprintf(msg + " [%v]", time.Now())
		}
		if traceMessageFlag & PRINT_PROC == PRINT_PROC {
			msg = fmt.Sprintf(msg + " <%s>", name)
		}
		if traceMessageFlag & PRINT_FILE == PRINT_FILE && traceMessageFlag & PRINT_LINE == PRINT_LINE {
			msg = fmt.Sprintf(msg + " (%s:%d)", sourceFile, line)
		} else if traceMessageFlag & PRINT_FILE == PRINT_FILE {
			msg = fmt.Sprintf(msg + " (%s)", sourceFile)
		} else if traceMessageFlag & PRINT_LINE == PRINT_LINE {
			msg = fmt.Sprintf(msg + " (%d)", line)
		}

		msg = fmt.Sprintf(msg + " - %v", variable)
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

		var msg string = "TRACE"
		if traceMessageFlag & PRINT_NONE == PRINT_NONE {
			msg = msg + ""
		}
		if traceMessageFlag & PRINT_TIME == PRINT_TIME {
			msg = fmt.Sprintf(msg + " [%v]", time.Now())
		}
		if traceMessageFlag & PRINT_PROC == PRINT_PROC {
			msg = fmt.Sprintf(msg + " <%s>", name)
		}
		if traceMessageFlag & PRINT_FILE == PRINT_FILE && traceMessageFlag & PRINT_LINE == PRINT_LINE {
			msg = fmt.Sprintf(msg + " (%s:%d)", sourceFile, line)
		} else if traceMessageFlag & PRINT_FILE == PRINT_FILE {
			msg = fmt.Sprintf(msg + " (%s)", sourceFile)
		} else if traceMessageFlag & PRINT_LINE == PRINT_LINE {
			msg = fmt.Sprintf(msg + " (%d)", line)
		}

		msg = fmt.Sprintf(msg + " - %v", umsg)
		return msg
	}
	return ""
}

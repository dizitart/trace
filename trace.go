package tracer

import (
	"fmt"
	"time"
	"runtime"
	"os"
	"path/filepath"
	"io"
)

// Use these configuration variables to set tracing parameters
var (
	// Flag to enable or disable the Tracing
	ENABLE_TRACE		bool 	= false
	// Flag to set the tracer to print in Stdout
	PRINT_TO_STDOUT		bool 	= true
	// Flag to set the tracer to print in Stdout
	PRINT_TO_FILE		bool 	= false
	// Trace log file
	TRACE_OUT_FILE		string
	// Trace message format
	TRACE_FORMAT		string	= "TRACE [%v] <%s> (%s:%d) - %v"
)

// Trace writes information about current source file, line number and input variable
func Trace(variable interface {}) {
	if ENABLE_TRACE {
		msg := createTraceMessage(variable)
		if PRINT_TO_STDOUT {
			fmt.Println(msg)
		}

		if PRINT_TO_FILE {
			if TRACE_OUT_FILE == "" {
				return
			} else {
				file, _ := os.OpenFile(TRACE_OUT_FILE, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
				defer file.Close()
				file.WriteString(msg + "\n")
			}
		}
	}
}

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

// Trace writes information about current source file, line number and input variable after
// formatting it according to a format specifier
func Tracef(format string, variables ...interface{}) {
	if ENABLE_TRACE {
		msg := createTraceMessagef(format, variables)
		if PRINT_TO_STDOUT {
			fmt.Println(msg)
		}

		if PRINT_TO_FILE {
			if TRACE_OUT_FILE == "" {
				return
			} else {
				file, _ := os.OpenFile(TRACE_OUT_FILE, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
				defer file.Close()
				file.WriteString(msg + "\n")
			}
		}
	}
}



func createTraceMessagef(format string, variables ...interface{}) string {
	pc, path, line, ok := runtime.Caller(1)
	if ok {
		fun := runtime.FuncForPC(pc)
		name := fun.Name()

		sourceFile := filepath.Base(path)
		umsg := fmt.Sprintf(format, variables)
		msg := fmt.Sprintf(TRACE_FORMAT, time.Now(), name, sourceFile, line, umsg)
		return msg
	}
	return ""
}

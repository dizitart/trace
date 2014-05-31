#Go Trace Utility

[![Build Status](https://travis-ci.org/dizitart/trace.svg?branch=master)](https://travis-ci.org/dizitart/trace)
[![Coverage Status](https://coveralls.io/repos/anidotnet/assert/badge.png?branch=master)](https://coveralls.io/r/anidotnet/assert?branch=master)
[![GoDoc](https://godoc.org/github.com/dizitart/trace?status.png)](https://godoc.org/github.com/dizitart/trace)

##Install

    $ go get github.com/dizitart/trace
    
##Use

```Go
package main

import (
	"fmt"
	"github.com/dizitart/trace"
	"time"
)

func main() {
	fmt.Println("Hello World")

	// Set up tracing
	trace.ENABLE_TRACE = true
	trace.PRINT_TO_FILE = true
	trace.TRACE_OUT_FILE = "/temp/trace.log"

	someString := "Hello World"
	// assert
	trace.Assert(someString == "hello world")
	// formatted assertion
	trace.Assertf(someString == "hello world", "Assertion Failed! at %v", time.Now())

	// write
	trace.Write(someString)
	// formatted write
	trace.Writef("Tracing at %v for value %s", time.Now(), someString)
}
```

##Doc
    https://godoc.org/github.com/dizitart/trace

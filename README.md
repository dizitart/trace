#Go Trace Utility

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


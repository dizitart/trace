package trace

import (
	"testing"
	"io/ioutil"
	"fmt"
	"runtime"
	"os"
	"strings"
	"strconv"
)

var tempFile *os.File

// CAUTION!!!! - Line sensitive test - Line = 20
func TestWrite(t *testing.T) {
	init_trace()
	defer close_trace()
	someVariable := "write-test"
	Write(someVariable)
	verify(t, someVariable, 20, "trace_test.go")
}

// CAUTION!!!! - Line sensitive test - Line = 29
func TestWritef(t *testing.T) {
	init_trace()
	defer close_trace()
	someVariable := "write-test"
	Writef("format-test %s", someVariable)
	finalmsg := fmt.Sprintf("format-test %s", someVariable)
	verify(t, finalmsg, 29, "trace_test.go")
}

// CAUTION!!!! - Line sensitive test - Line = 38
func TestAssert(t *testing.T) {
	init_trace()
	defer close_trace()
	Assert(2 == 3)
	verify(t, "Assertion Failed!", 38, "trace_test.go")
}

// CAUTION!!!! - Line sensitive test - Line = 46
func TestAssertf(t *testing.T) {
	init_trace()
	defer close_trace()
	Assertf("False %v", 2 == 3)
	verify(t, "False Assertion Failed!", 46, "trace_test.go")
}

func verify(t *testing.T, message string, line int, file string) {
	msg, err := ioutil.ReadAll(tempFile)
	if err != nil {
		assert(t, false, 1)
	}
	if !strings.Contains(string(msg), message) {
		assert(t, false, 1)
	}

	fileLine := file + ":" + strconv.Itoa(line)
	if !strings.Contains(string(msg), fileLine) {
		assert(t, false, 1)
	}
}

func init_trace() {
	ENABLE_TRACE = true
	var err error
	tempFile, err = ioutil.TempFile("", "trace_test")
	if err != nil {
		fmt.Println("Couldnot create temp file.")
		runtime.Goexit()
	}
	TRACE_OUT_FILE = tempFile.Name()
	PRINT_TO_FILE = true
}

func close_trace() {
	tempFile.Close()
	os.Remove(TRACE_OUT_FILE)
	PRINT_TO_FILE = false
	ENABLE_TRACE = false
}

func assert(t *testing.T, result bool, cd int) {
	if !result {
		_, file, line, _ := runtime.Caller(cd + 1)
		t.Errorf("%s:%d", file, line)
		t.FailNow()
	}
}

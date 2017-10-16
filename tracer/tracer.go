package tracer

import (
	"fmt"
	"io"
	"os"
)

const (
	k_FATAL	= "FATAL: "
	k_ERROR = "E: "
	k_WARN  = "W: "
	k_INFO  = "I: "
	k_TRACE	= "[TRACE%d] "
)

var (
	threshold	uint = 0	
	x1	io.Writer
	x2	io.Writer
	
)

func init() {
	x1 = os.Stdout
	x2 = os.Stderr
}


func SetLevel(x int) uint {
	o := threshold
	if x < 0 {
		threshold = 0
	} else {
		threshold = uint(x)
	}
	return o
}

func SetupTrace(out io.Writer, n uint) {
	if nil != out {
		x1 = out
	}
	threshold = n
}

func Fatal(x string) {
	// FATAL: ...
	fmt.Fprintf(x2, k_FATAL)
	fmt.Fprintln(x2, x)
	os.Exit(255)
}

func FatalErr(e error) {
	// FATAL: ...
	fmt.Fprintf(x2, k_FATAL)
	fmt.Fprintln(x2, e.Error())
	fmt.Fprintln(x2, "")
	os.Exit(255)
}



func Error(x string) {
	fmt.Fprintf(x2, k_ERROR)
	fmt.Fprintln(x2, x)
}

func Warn(x string) {
	fmt.Fprintf(x2, k_WARN)
	fmt.Fprintln(x2, x)
}

func Info(x string) {
	fmt.Fprintf(x2, k_INFO)
	fmt.Fprintln(x2, x)
}




func Trace(n uint, s string) {
	if n > threshold {
		return
	}
	fmt.Fprintf(x1, k_TRACE, n)
	fmt.Fprintln(x1, s)
}



func Tracef(n uint, s string, x ...interface{}) {
	if n > threshold {
		return
	}
	fmt.Fprintf(x1, k_TRACE, n)
	fmt.Fprintf(x1, s, x...)
}

func TraceV(n uint, x ...interface{}) {
	if n > threshold {
		return
	}
	fmt.Fprintf(x1, k_TRACE, n)
	fmt.Fprintln(x1, x...)
}

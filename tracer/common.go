package tracer

import (
	"io"
	"os"
)

const (
	k_FATAL	= "FATAL: "
	k_ERROR = "E: "
	k_WARN  = "W: "
	k_INFO  = "I: "
	k_NOTICE= "N: "
	k_TRACE	= "[TRACE%d] "
	k_DEBUG = "DEBUG: "
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

func SetOutput(out io.Writer) {
	if nil != out {
		x2 = out
	}
}

func _puts(w io.Writer, x string) (err error) {
	_,err = io.WriteString(w,x)
	return
}

func _putln(w io.Writer, x string) (err error) {
	_,err = io.WriteString(w,x)
	w.Write([]byte("\n"))
	return
}

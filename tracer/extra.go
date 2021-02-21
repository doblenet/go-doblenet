package tracer

import (
	"fmt"
)


func FatalV(z ...interface{}) {
	// FATAL: ...
	_puts(x2, k_FATAL)
	fmt.Fprintln(x2, z...)
}

func ErrorV(z ...interface{}) {
	_puts(x2, k_ERROR)
	fmt.Fprintln(x2, z...)
}

func WarnV(z ...interface{}) {
	_puts(x2, k_WARN)
	fmt.Fprintln(x2, z...)
}

func InfoV(z ...interface{}) {
	_puts(x2, k_INFO)
	fmt.Fprintln(x2, z...)
}

func NoticeV(z ...interface{}) {
	_puts(x2, k_NOTICE)
	fmt.Fprintln(x2, z...)
}


func Putf(x string,z ...interface{}) {
	fmt.Fprintf(x1, x, z...)
}

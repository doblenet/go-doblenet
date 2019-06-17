package tracer

import (
	"os"
)


func FatalErr(e error) {
	// FATAL: ...
	_puts(x2, k_FATAL)
	_putln(x2, e.Error())
	os.Exit(3)
}

func FatalExit(n int, x string) {
	_puts(x2, k_FATAL)
	_putln(x2, x)
	os.Exit(n)
}


func Fatal(x string) {
	// FATAL: ...
	_puts(x2, k_FATAL)
	_putln(x2, x)
}

func Error(x string) {
	_puts(x2, k_ERROR)
	_putln(x2, x)
}

func Warn(x string) {
	_puts(x2, k_WARN)
	_putln(x2, x)
}

func Info(x string) {
	_puts(x2, k_INFO)
	_putln(x2, x)
}

func Notice(x string) {
	_puts(x2, k_NOTICE)
	_putln(x2, x)
}

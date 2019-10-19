package tracer

import (
	"os"
)


func Puts(x string) {
	_puts(os.Stdout,x)
}

func Putln(x string) {
	_putln(os.Stdout,x)
}

func PutV(s string, vals ...string) {
	buf := make([]byte,0,len(s)*2)
	buf=append(buf,[]byte(s)...)
	for _,x := range vals {
		buf=append(buf,[]byte(x)...)
	}
	_putln(os.Stdout,string(buf))
}


func TraceOver(level uint) bool {
	return (threshold <= level)
}

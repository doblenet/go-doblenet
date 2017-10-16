package err


import (
	"fmt"
)



type exitErr struct {
	data	string	//interface{}
	code	int
}


type ExitCoder interface {
	error
	ExitCode()	int
}



func ExitError(exitCode int, info error) ExitCoder {
	return &exitErr{
		code: exitCode,
		data: info.Error(),
	}
}

func ExitErrorv(exitCode int, msg ...interface{}) ExitCoder {
	return &exitErr{
		code: exitCode,
		data: fmt.Sprint(msg...),
	}
}

func ExitErrorf(exitCode int, format string, msg ...interface{}) ExitCoder {
	return &exitErr{
		code: exitCode,
		data: fmt.Sprintf(format,msg...),
	}
}

// Error returns the string message, fulfilling the interface required by
// `error`
func (ee *exitErr) Error() string {
	return ee.data		// fmt.Sprintf("%v", ee.data)
}

// ExitCode returns the exit code, fulfilling the interface required by
// `ExitCoder`
func (ee *exitErr) ExitCode() int {
	return ee.code
}

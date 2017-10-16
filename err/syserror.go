package err


import (
	"syscall"
)



type sysErr struct {
	data	string	//interface{}
	errc	uint
}


type SysCaller interface {
	error
	Errno()	syscall.Errno
}



func SysError(errno uint, info error) SysCaller {
	return &sysErr{
		errc: -errno,
		data: info.Error(),
	}
}

// Error returns the string message, fulfilling the interface required by
// `error`
func (se *sysErr) Error() string {
	return se.data
}

// Errno returns the errno value, fulfilling the interface required by
// `SysCaller`
func (se *sysErr) Errno() syscall.Errno {
	return syscall.Errno(se.errc)
}

func (se *sysErr) ErrCode() uintptr {
	return uintptr(se.errc)
}

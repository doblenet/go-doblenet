package err


import (
    "os"
    "syscall"
)

func ConvertErrno(err error) syscall.Errno {

    if e, ok := err.(*os.PathError); ok {
    
        c := e.Err
        if r,g := c.(syscall.Errno); g {
            return r
        }
    }			

    return 0
}

func CheckSysErr(err error, val uintptr) bool {

    // Check os.PathError first
    if e, ok := err.(*os.PathError); ok {
       
        c := e.Err
	if r,g := c.(syscall.Errno); g {
	    return r  == syscall.Errno(val)
	}
	// Not a *os.PathError
    }
    
    
    // Check if generic syscall.Errno
    if e, ok := err.(syscall.Errno); ok {
        return e == syscall.Errno(val)
    }

    return false
}

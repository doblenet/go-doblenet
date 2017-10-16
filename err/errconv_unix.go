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

func CheckErrno(err error, val syscall.Errno) bool {

    // Check os.PathError first
    if e, ok := err.(*os.PathError); ok {
       
        c := e.Err
	if r,g := c.(syscall.Errno); g {
	    return r  == val
	}
	// Not a *os.PathError
    }
    
    
    // Check if generic syscall.Errno
    if e, ok := err.(syscall.Errno); ok {
        return e == val
    }

    return false
}

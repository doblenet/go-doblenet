package err

import (
    "os/exec"
    "syscall"
)


func ExecErrorConv(err error) (r int, msg string) {

    r = 192; msg=""
    if ee, ok := err.(*exec.ExitError); ok {
        msg = ee.Error()
        r = 86
        if status, ok := ee.Sys().(syscall.WaitStatus); ok {
            r = status.ExitStatus()
        }
    }
    return
}

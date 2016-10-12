package eflags

import (
	"github.com/spf13/pflag"	// integration with Cobra / pflag
)

var (
	cmdLine = pflag.CommandLine
)

//func setVarP() {
//	return cmdLine.VarP()
//

//go:build aix || darwin || dragonfly || freebsd || js || linux || netbsd || openbsd || solaris
// +build aix darwin dragonfly freebsd js linux netbsd openbsd solaris

package main

import (
	"os"
	"syscall"
)

var signalsToIgnore = []os.Signal{os.Interrupt, syscall.SIGQUIT}

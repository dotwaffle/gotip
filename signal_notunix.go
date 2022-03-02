//go:build plan9 || windows
// +build plan9 windows

package main

import (
	"os"
)

var signalsToIgnore = []os.Signal{os.Interrupt}

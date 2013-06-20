// +build !linux,!darwin

package loadavg

import (
	"fmt"
	"runtime"
)

func loadAvg() ([3]float64, error) {
	return [...]float64{-1, -1, -1}, fmt.Errorf("loadavg: unsupported platform %q", runtime.GOOS)
}

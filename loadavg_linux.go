package loadavg

import (
	"fmt"
	"os"
)

const procfile = "/proc/loadavg"

func loadAvg() ([3]float64, error) {
	f, err := os.Open(procfile)
	if err != nil {
		return [...]float64{-1, -1, -1}, fmt.Errorf("loadavg: unable to open procfile %q: %v", procfile, err)
	}
	defer f.Close()
	var loadavg [3]float64
	if n, err := fmt.Fscanf(f, "%f %f %f", &loadavg[0], &loadavg[1], &loadavg[2]); n != 3 || err != nil {
		return [...]float64{-1, -1, -1}, fmt.Errorf("loadavg: unable to read loadavg: %v", err)
	}
	return loadavg, nil
}

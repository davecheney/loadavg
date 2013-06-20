package loadavg

const (
	ONE_MIN  = 0
	FIVE_MIN = iota
	FIFTEEN_MIN
)

// LoadAvg returns the traditional 1, 5, and 15 min load averages.
func LoadAvg() ([3]float64, error) { return loadAvg() }

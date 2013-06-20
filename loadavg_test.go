package loadavg

import (
	"fmt"
	"log"
	"testing"
)

func TestLoadAvg(t *testing.T) {
	loadavg, err := LoadAvg()
	if err != nil {
		t.Fatal(err)
	}
	for _, l := range loadavg[:] {
		if l < 0 {
			t.Errorf("expected loadavg >= 0, got %v", l)
		}
	}
	t.Logf("loadavg: %2.2f, %2.2f, %2.2f", loadavg[ONE_MIN], loadavg[FIVE_MIN], loadavg[FIFTEEN_MIN])
}

func ExampleLoadAvg() {
	loadavg, err := LoadAvg()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("loadavg: %2.2f, %2.2f, %2.2f", loadavg[ONE_MIN], loadavg[FIVE_MIN], loadavg[FIFTEEN_MIN])
}

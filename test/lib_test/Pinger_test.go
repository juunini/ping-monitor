package lib

import (
	"ping-monitor/lib"
	"testing"
)

func TestPinger(t *testing.T) {
	pinger := lib.NewPinger()
	result := pinger.Do()

	for _, r := range result {
		t.Log("\n", r.Message, "\n", r.Error)
	}
}
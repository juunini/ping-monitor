package monitor_test

import (
	"ping-monitor/lib"
	"ping-monitor/monitor"
	"testing"
	"time"
)

func TestMonitoring(t *testing.T) {
	monitor.Monitoring("./test.json")

	for {
		lib.Logging("monitoring...")
		time.Sleep(time.Minute)
	}
}

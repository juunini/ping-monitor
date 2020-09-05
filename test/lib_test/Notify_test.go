package lib

import (
	"github.com/joho/godotenv"
	"ping-monitor/lib"
	"testing"
)

func TestNotify(t *testing.T) {
	_ = godotenv.Load()
	lib.Notify("test", "error test")
	lib.Notify("test1", "error test1")
	lib.Notify("test2", "error test2")
}

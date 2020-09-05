package lib

import (
	"github.com/juunini/simple-go-line-notify/notify"
	"os"
)

func Notify(name, err string) {
	Logging("[WARN]", name, "-", err)

	accessKey := os.Getenv("LINE_ACCESS_KEY")
	_ = notify.SendText(accessKey, name+" - "+err)
}

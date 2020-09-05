package monitor

import (
	"fmt"
	"ping-monitor/lib"
	"time"
)

func Monitoring(serverListFile string) {
	var conf config
	if err := lib.ParseJSON(serverListFile, &conf); err != nil {
		panic(err)
	}

	for _, server := range conf {
		go infiniteLoopPing(server)
	}
}

func infiniteLoopPing(server serverInfo) {
	pinger := lib.Pinger{
		Host:    server.Host,
		Port:    server.Port,
		Network: server.Network,
		Repeat:  server.Repeat,
		Timeout: time.Duration(server.Timeout) * time.Millisecond,
	}

	for {
		go ping(pinger, server)

		time.Sleep(time.Duration(server.Peroid) * time.Millisecond)
	}
}

func ping(pinger lib.Pinger, server serverInfo) {
	var (
		err        error
		logMessage string
	)
	result := pinger.Do()

	for _, row := range result {
		if row.Error != nil {
			err = row.Error
			break
		}
		logMessage += fmt.Sprintf("\n\t[%s] %s", server.Name, row.Message)
	}

	if err != nil {
		lib.Notify(server.Name, err.Error())
	} else {
		lib.Logging(logMessage)
	}
}

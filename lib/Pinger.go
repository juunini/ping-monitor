package lib

import (
	"fmt"
	"net"
	"time"
)

type Pinger struct {
	Host    string        `default:"127.0.0.1"`
	Port    uint16        `default:"80"`
	Network string        `default:"tcp"`
	Repeat  uint          `default:"3"`
	Timeout time.Duration `default:"1s"`
}

type PingerResult struct {
	Message string
	Error   error
}

func NewPinger() Pinger {
	return Pinger{
		"127.0.0.1",
		80,
		"tcp",
		3,
		1 * time.Second,
	}
}

func (p *Pinger) Do() (result []PingerResult) {
	address := fmt.Sprintf("%s:%d", p.Host, p.Port)

	var i uint
	for i = 0; i < p.Repeat; i++ {
		duration, err := ping(
			p.Network,
			address,
			p.Timeout,
		)
		message := makePingMessage(address, p.Network)
		if err != nil {
			result = append(result, PingerResult{
				Message: message + " - timeout",
				Error:   err,
			})
		} else {
			result = append(result, PingerResult{
				Message: fmt.Sprintf("%s - port is open - time=%s", message, duration),
			})
		}
	}
	return
}

func ping(network string, address string, timeout time.Duration) (duration time.Duration, err error) {
	start := time.Now()

	p, err := net.DialTimeout(network, address, timeout)
	if err != nil {
		return
	}
	defer p.Close()

	duration = time.Since(start)
	return
}

func makePingMessage(address, network string) string {
	return fmt.Sprintf(
		"[%s] Probing %s/%s",
		time.Now().Format(time.RFC3339),
		address,
		network,
	)
}

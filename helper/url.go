package helper

import (
	"fmt"
	"github.com/tatsushid/go-fastping"
	"net"
	"strings"
	"time"
)

type Url struct {
	ip   string
	isOk bool
	port string
	time time.Duration
}

func createUrl(s string) *Url {
	startInd := strings.Index(s, ":")
	endInd := startInd + 1
	if startInd == -1 {
		startInd = len(s)
		endInd = startInd
	}
	return &Url{ip: s[:startInd], port: s[endInd:]}
}

func (url *Url) ping() {
	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", url.ip)
	if err != nil {
		return
	}
	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, duration time.Duration) {
		if duration <= pingTimeMax {
			url.isOk = true
			url.time = duration
		}
	}
	errFatalNotification(p.Run())
}

func pingUrls() {
	stringNotification("Start pinging...")
	for i, s := range addresses {
		stringNotification(fmt.Sprintf("Url %d/%d...", i + 1, len(addresses)))
		url := createUrl(s)
		url.ping()
		if url.isOk {
			urls = append(urls, *createUrl(s))
		}
	}
	stringNotification("Successfully completed.")
}

package main

import (
	"github.com/sirupsen/logrus"
	"net"
)

func parseIP(s string) net.IP {
	ip := net.ParseIP(s)
	if ip == nil {
		return nil
	}
	return ip
}

func main() {
	var ips []string = make([]string, 0)
	ips = append(ips, "131.123.234. 234")
	ips = append(ips, "")
	for _, ip := range ips {
		if len(ip) > 0 {
			ip22 := parseIP(ip)
			logrus.Info(ip22)
		}
	}
}

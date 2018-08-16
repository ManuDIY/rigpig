package serviceDiscovery

import "time"

type Consul struct {
	Name   string
	Uptime time.Duration
}

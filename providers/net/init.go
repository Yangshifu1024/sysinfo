package net

import (
	"github.com/shirou/gopsutil/net"
)

type NetInterface struct {
	Name string
	Addr string
}

var (
	NetInterfaces []NetInterface
    NetInterfaceCount int
)

func init() {
	NetInterfaces = make([]NetInterface, 0)

	stats, err := net.NetInterfaces()
	if err != nil {
		panic(err)
	}
	for _, stat := range stats {
		if len(stat.Name) != 0 && len(stat.Addrs) > 0 {
			inter := NetInterface{
				Name: stat.Name,
				Addr: stat.Addrs[0].Addr,
			}
			NetInterfaces = append(NetInterfaces, inter)
		}
	}
    NetInterfaceCount = len(NetInterfaces)
}

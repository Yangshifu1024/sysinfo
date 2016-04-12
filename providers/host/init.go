package host

import (
	"time"

	"github.com/shirou/gopsutil/host"

	"github.com/iFrankYang/iSystem/formats"
)

var (
	HostName             string
	Procs                uint64
	OS                   string
	Platform             string
	PlatformFamily       string
	PlatformVersion      string
	VirtualizationSystem string
	VirtualizationRole   string
	Users                int
	BootTime             string
	Uptime               string
)

func init() {
	hostInfo, err := host.Info()
	if err != nil {
		panic(err)
	}
	users, err := host.Users()
	if err != nil {
		panic(err)
	}

	HostName = hostInfo.Hostname
	Procs = hostInfo.Procs
	OS = hostInfo.OS
	Platform = hostInfo.Platform
	PlatformFamily = hostInfo.PlatformFamily
	PlatformVersion = hostInfo.PlatformVersion
	VirtualizationSystem = hostInfo.VirtualizationSystem
	VirtualizationRole = hostInfo.VirtualizationRole
	Users = len(users)
	BootTime = time.Unix(int64(hostInfo.BootTime), 0).Format("2006-01-02 15:04:05")
	Uptime = formats.Success((time.Duration(hostInfo.Uptime) * time.Second).String())
}

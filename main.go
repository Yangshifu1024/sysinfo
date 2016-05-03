package main

import (
	"fmt"
	"os"
	"time"

	"github.com/codegangsta/cli"
	"github.com/gosuri/uitable"

	"github.com/codebear4/sysinfo/formats"
	"github.com/codebear4/sysinfo/providers/cpu"
	"github.com/codebear4/sysinfo/providers/host"
	"github.com/codebear4/sysinfo/providers/load"
	"github.com/codebear4/sysinfo/providers/mem"
	"github.com/codebear4/sysinfo/providers/net"
)

var VERSION string

func main() {
	app := cli.NewApp()
	app.Name = "iSystem"
	app.Usage = "Show Human-Readable system info of *nix."
	app.Authors = []cli.Author{
		{
			Name:  "Frank Yang",
			Email: "codebear4@gmail.com",
		},
	}
	app.Copyright = "Codebear (c) 2016"
	app.Version = VERSION

	var noCPU bool = false

	app.Flags = []cli.Flag{
		cli.BoolFlag{Name: "no-cpu", Usage: "Hide CPU information.", Destination: &noCPU},
	}

	app.Action = func(c *cli.Context) error {
		infoTable := uitable.New()
		infoTable.MaxColWidth = 30
		infoTable.Wrap = true
		fmt.Println(formats.Title("System information as of ", time.Now().String()))

		if !noCPU {
			infoTable.AddRow("CPU Model:", cpu.ModelName, "Cores:", cpu.Cores)
		}
		infoTable.AddRow("Platform:", host.Platform, "PlatformVersion:", host.PlatformVersion)
		if host.VirtualizationSystem != "" {
			infoTable.AddRow("Virtualization:", host.VirtualizationSystem, "Role:", host.VirtualizationRole)
		}
		infoTable.AddRow("Hostname:", host.HostName, "Users:", host.Users)
		infoTable.AddRow("Boot time:", host.BootTime, "Uptime:", host.Uptime)
		infoTable.AddRow("Processes:", host.Procs, "System load:", load.Load)
		infoTable.AddRow("Memory total:", mem.Total, "Free:", mem.Free)
		infoTable.AddRow("Memory Used:", mem.Used, "Usage:", mem.Usage)
		if mem.SwapTotal != "0" {
			infoTable.AddRow("Swap total:", mem.SwapTotal, "Free:", mem.SwapFree)
			infoTable.AddRow("Swap Used:", mem.SwapUsed, "Usage:", mem.SwapUsage)
		}
		for i := 0; i <= net.NetInterfaceCount; i += 2 {
			if i >= net.NetInterfaceCount {
				break
			}
			cInter := net.NetInterfaces[i]
			if i < net.NetInterfaceCount-1 {
				nInter := net.NetInterfaces[i+1]
				infoTable.AddRow(cInter.Name+":", cInter.Addr, nInter.Name+":", nInter.Addr)
			} else {
				infoTable.AddRow(cInter.Name+":", cInter.Addr)
			}
		}
		fmt.Println(infoTable)
		return nil
	}

	app.Run(os.Args)
}

package main

import (
	"fmt"
	"os"

	"iSystem/providers"
	"github.com/codegangsta/cli"
)

var VERSION string

func main() {
	app := cli.NewApp()
	app.Name = "xInfo"
	app.Usage = "Show Human-Readable system info of *nix."
	app.Authors = []cli.Author{
		{
			Name:  "Frank Yang",
			Email: "codebear4@gmail.com",
		},
	}
	app.Copyright = "Codebear (c) 2016"
	app.Version = VERSION

	app.Action = func(c *cli.Context) {
		cpuStats := providers.CpuStats()
		fmt.Println(cpuStats)
		memStats := providers.MemStats()
		fmt.Println(memStats)
		userStats := providers.UserStats()
		fmt.Println(userStats)
	}

	app.Run(os.Args)
}

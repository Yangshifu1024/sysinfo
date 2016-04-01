package load

import (
	"github.com/shirou/gopsutil/load"
)

var (
	Load1  float64
	Load5  float64
	Load15 float64
)

func init() {
	loadInfo, err := load.LoadAvg()
	if err != nil {
		panic(err)
	}

	Load1 = loadInfo.Load1
	Load5 = loadInfo.Load5
	Load15 = loadInfo.Load15
}

package cpu

import (
	"github.com/shirou/gopsutil/cpu"
)

var (
    ModelName string
    Cores int32
)

func init() {
	cpuStats, err := cpu.CPUInfo()
	if err != nil {
		panic(err)
	}
    cpuStat := cpuStats[0]
    ModelName = cpuStat.ModelName
    Cores = cpuStat.Cores
}
